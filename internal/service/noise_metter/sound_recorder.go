package noise_metter

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type RecordTask struct {
	StartTime time.Time
	Duration  time.Duration
}

func (s *Service) addRecordTask(startTime time.Time, duration time.Duration) {
	s.recordTasks <- &RecordTask{
		StartTime: startTime,
		Duration:  duration,
	}
}

func (s *Service) bgFetchRecordTasks() {
	for {
		select {
		case <-s.ctx.Done():
			return
		case task := <-s.recordTasks:
			s.RecordSoundWrapper(task)
		}
	}
}

func (s *Service) RecordSoundWrapper(task *RecordTask) {
	// wait till now will be after task.StartTime+task.Duration + 4 seconds just in case
	endTime := task.StartTime.Add(task.Duration).Add(4 * time.Second)
	for {
		if time.Now().After(endTime) {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err := s.RecordSound(task); err != nil {
		s.log.Error("failed to record sound: ", err)
	}
}

func (s *Service) RecordSound(task *RecordTask) error {
	url := fmt.Sprintf("ws://%s/api/stream2/", s.conf.RemoteHost)
	conn, _, err := websocket.DefaultDialer.DialContext(s.ctx, url, nil)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	defer conn.Close()

	// Expect initial "Password:\n"
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	if err = conn.WriteMessage(websocket.TextMessage, []byte("admin")); err != nil {
		return fmt.Errorf("send session ID failed: %w", err)
	}
	// identification
	q := fmt.Sprintf("Audio %d, %d\n", task.StartTime.UnixMilli(), int64(task.Duration))
	if err = conn.WriteMessage(websocket.TextMessage, []byte(q)); err != nil {
		return fmt.Errorf("send session ID failed: %w", err)
	}

	result := make([]byte, 0, 1_000_000)
	for {
		_, msg, err = conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		msgStr := string(msg)
		if strings.Contains(msgStr, "XL3 Streaming API Text") {
			continue
		}
		if strings.HasPrefix(msgStr, "4;2") {
			break // end of stream
		}

		dataChunked := strings.Split(msgStr, ";")
		if len(dataChunked) == 3 {
			dataD, errD := base64.StdEncoding.DecodeString(dataChunked[2])
			if errD != nil {
				return fmt.Errorf("base64 decode error: %w", errD)
			}
			result = append(result, dataD...)
		}
	}
	if err = s.repo.DumpNoiseAudioRaw(result); err != nil {
		return fmt.Errorf("write file error: %w", err)
	}
	return nil
}
