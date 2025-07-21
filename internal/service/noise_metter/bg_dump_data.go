package noise_metter

import (
	"noize_metter/internal/logger"
	"time"
)

var (
	dumpDataDuration = 1 * time.Minute // Duration for saving candidates
)

func (s *Service) bgDumpData() {
	ticker := time.NewTicker(dumpDataDuration)
	defer ticker.Stop()
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			s.dumpData()
		}
	}
}

func (s *Service) dumpData() {
	data := s.items.LoadAndErase()
	if len(data) == 0 {
		return
	}
	s.log.Info("dumping noise data", logger.WithInt("items", len(data)))
	if err := s.repo.DumpNoiseRawData(data); err != nil {
		s.log.Error("error saving noise data", err)
	}
}
