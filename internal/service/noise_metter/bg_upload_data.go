package noise_metter

import "time"

var (
	uploadDataDuration = 1 * time.Minute // Duration for uploading data
)

func (s *Service) bgUploadData() {
	ticker := time.NewTicker(uploadDataDuration)
	defer ticker.Stop()
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			s.uploadData()
		}
	}
}

func (s *Service) uploadData() {}
