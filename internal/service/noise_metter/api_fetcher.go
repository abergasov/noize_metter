package noise_metter

import (
	"encoding/json"
	"fmt"
	"log"
	"noize_metter/internal/entities"
	"noize_metter/internal/utils"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func (s *Service) FetchAPIsS() error {
	url := fmt.Sprintf("ws://%s/api/stream1/", s.conf.RemoteHost)
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
	if len(msg) > 0 {
	}
	if err = conn.WriteMessage(websocket.TextMessage, []byte("admin")); err != nil {
		return fmt.Errorf("send session ID failed: %w", err)
	}
	// identification
	_, ident, err := conn.ReadMessage()
	if err != nil {
		log.Fatalf("recv ident: %v", err)
	}
	fmt.Println("11111", strings.TrimSpace(string(ident)))

	now := (time.Now().Unix() - 60) * 1000
	q := fmt.Sprintf("WEATHER %d\n", now)
	q = fmt.Sprintf(`SPLLOG %d, "LCEQ_G1-LAEQ_G1"\n`, now)
	q = fmt.Sprintf(`SPLLOG %d, "LAEQ LAFMAX"\n`, now)
	q = fmt.Sprintf(`SPLLOG %d, "LAEQ LAFMAX"\n`, now)
	if err = conn.WriteMessage(websocket.TextMessage, []byte(q)); err != nil {
		log.Fatalf("send query: %v", err)
	}
	//if err = conn.WriteMessage(websocket.TextMessage, []byte(`Audio 1690547362000, 3`)); err != nil {
	//	return fmt.Errorf("send START command failed: %w", err)
	//}

	type streamData struct {
		Data struct {
			Timer  string    `json:"timer"`
			Field2 []float64 `json:"123"`
		} `json:"data"`
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		if strings.Contains(string(msg), "settings") {
			continue
		}
		continue
		var data streamData
		if err = json.Unmarshal(msg, &data); err != nil {
			return fmt.Errorf("unmarshal data: %w", err)
		}
		if len(data.Data.Field2) != 5 {
			return fmt.Errorf("expected 5 noise measures, got %d", len(data.Data.Field2))
		}
		tms := time.Now()
		s.items.Add(entities.NoiseMeasures{
			Timestamp:    tms,
			TimestampNum: utils.TimeToDayIntNum(tms),

			LAeqDT:  data.Data.Field2[0],
			LAf:     data.Data.Field2[1],
			LCPK:    data.Data.Field2[2],
			LAeqG10: data.Data.Field2[3],
			LAeqG5:  data.Data.Field2[4],
		})
	}
}

func (s *Service) FetchAPIs() error {
	url := fmt.Sprintf("ws://%s/api/stream1/", s.conf.RemoteHost)
	conn, _, err := websocket.DefaultDialer.DialContext(s.ctx, url, nil)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	defer conn.Close()

	// Expect initial "Password:\n"
	//_, msg, err := conn.ReadMessage()
	//if err != nil {
	//	return fmt.Errorf("read: %w", err)
	//}
	if err = conn.WriteMessage(websocket.TextMessage, []byte("admin")); err != nil {
		return fmt.Errorf("send session ID failed: %w", err)
	}
	// identification
	_, ident, err := conn.ReadMessage()
	if err != nil {
		log.Fatalf("recv ident: %v", err)
	}
	fmt.Println(strings.TrimSpace(string(ident)))

	now := (time.Now().Unix() - 60) * 1000
	q := fmt.Sprintf("WEATHER %d\n", now)
	q = fmt.Sprintf("SOH\n")
	if err = conn.WriteMessage(websocket.TextMessage, []byte(q)); err != nil {
		log.Fatalf("send query: %v", err)
	}
	//if err = conn.WriteMessage(websocket.TextMessage, []byte(`Audio 1690547362000, 3`)); err != nil {
	//	return fmt.Errorf("send START command failed: %w", err)
	//}

	type streamData struct {
		Data struct {
			Timer  string    `json:"timer"`
			Field2 []float64 `json:"123"`
		} `json:"data"`
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		if strings.Contains(string(msg), "settings") {
			continue
		}
		continue
		var data streamData
		if err = json.Unmarshal(msg, &data); err != nil {
			return fmt.Errorf("unmarshal data: %w", err)
		}
		if len(data.Data.Field2) != 5 {
			return fmt.Errorf("expected 5 noise measures, got %d", len(data.Data.Field2))
		}
		tms := time.Now()
		s.items.Add(entities.NoiseMeasures{
			Timestamp:    tms,
			TimestampNum: utils.TimeToDayIntNum(tms),

			LAeqDT:  data.Data.Field2[0],
			LAf:     data.Data.Field2[1],
			LCPK:    data.Data.Field2[2],
			LAeqG10: data.Data.Field2[3],
			LAeqG5:  data.Data.Field2[4],
		})
	}
}

func (s *Service) FetchAPIsSensor() error {
	url := fmt.Sprintf("ws://%s/api/live/", s.conf.RemoteHost)
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
	if err = conn.WriteMessage(websocket.TextMessage, []byte("START 1234")); err != nil {
		return fmt.Errorf("send START command failed: %w", err)
	}

	type streamData struct {
		Data struct {
			Timer  string    `json:"timer"`
			Field2 []float64 `json:"123"`
		} `json:"data"`
	}
	for {
		_, msg, err = conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		if strings.Contains(string(msg), "settings") {
			continue
		}
		var data streamData
		if err = json.Unmarshal(msg, &data); err != nil {
			return fmt.Errorf("unmarshal data: %w", err)
		}
		if len(data.Data.Field2) != 5 {
			return fmt.Errorf("expected 5 noise measures, got %d", len(data.Data.Field2))
		}
		tms := time.Now()
		s.items.Add(entities.NoiseMeasures{
			Timestamp:    tms,
			TimestampNum: utils.TimeToDayIntNum(tms),

			LAeqDT:  data.Data.Field2[0],
			LAf:     data.Data.Field2[1],
			LCPK:    data.Data.Field2[2],
			LAeqG10: data.Data.Field2[3],
			LAeqG5:  data.Data.Field2[4],
		})
	}
}
