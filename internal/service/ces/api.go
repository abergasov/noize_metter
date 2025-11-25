package ces

import (
	"context"
	"fmt"
	"net/http"
	"noize_metter/internal/entities"
	"noize_metter/internal/utils/requests"
	"sync"
	"time"
)

func (s *Service) getToken(baseURL string) string {
	s.tokenMU.Lock()
	token, _ := s.token[baseURL]
	s.tokenMU.Unlock()
	return token
}

// GetAllMegaBoxes returns all mega boxes from all servers
// IMDC -> MegaboxID -> MegaBox
func (s *Service) GetAllMegaBoxes(ctx context.Context) ([]entities.MegaBox, error) {
	var (
		resultIMDC1 = make([]entities.MegaBox, 0, 4)
		errIMDC1    error
		resultIMDC2 = make([]entities.MegaBox, 0, 4)
		errIMDC2    error
		wg          sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, host := range URLListIMDC1 {
			resultIMDC1, errIMDC1 = s.getAllMegaBoxes(ctx, host)
			if len(resultIMDC1) > 0 {
				for i := range resultIMDC1 {
					resultIMDC1[i].IMDC = 1
				}
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for _, host := range URLListIMDC2 {
			resultIMDC2, errIMDC2 = s.getAllMegaBoxes(ctx, host)
			if len(resultIMDC2) > 0 {
				for i := range resultIMDC2 {
					resultIMDC2[i].IMDC = 2
				}
				break
			}
		}
	}()
	wg.Wait()

	if errIMDC1 != nil && len(resultIMDC1) == 0 {
		resultIMDC1 = []entities.MegaBox{
			{MegaBoxID: 1, IMDC: 1, Error: errIMDC1.Error()},
			{MegaBoxID: 2, IMDC: 1, Error: errIMDC1.Error()},
			{MegaBoxID: 3, IMDC: 1, Error: errIMDC1.Error()},
			{MegaBoxID: 4, IMDC: 1, Error: errIMDC1.Error()},
		}
	}
	if errIMDC2 != nil && len(resultIMDC2) == 0 {
		resultIMDC2 = []entities.MegaBox{
			{MegaBoxID: 1, IMDC: 2, Error: errIMDC2.Error()},
			{MegaBoxID: 2, IMDC: 2, Error: errIMDC2.Error()},
			{MegaBoxID: 3, IMDC: 2, Error: errIMDC2.Error()},
			{MegaBoxID: 4, IMDC: 2, Error: errIMDC2.Error()},
		}
	}
	return append(resultIMDC1, resultIMDC2...), nil
}

func (s *Service) getAllMegaBoxes(ctx context.Context, baseURL string) ([]entities.MegaBox, error) {
	token := s.getToken(baseURL)
	result, code, err := requests.GetCurl[[]entities.MegaBox](ctx,
		fmt.Sprintf("%s/api/megabox/getAllMegaboxes", baseURL),
		map[string]string{"auth-token": token},
		requests.WithSkipTLSVerification(true),
		requests.WithRequestTimeout(10*time.Minute),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get all mega boxes: %w", err)
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("wrong code in response of megaboxes: %d", code)
	}
	return result, nil
}

func (s *Service) GetAllTanks(ctx context.Context) ([]entities.Tank, error) {
	var (
		resultIMDC1 = make([]entities.Tank, 0, 4)
		errIMDC1    error
		resultIMDC2 = make([]entities.Tank, 0, 4)
		errIMDC2    error
		wg          sync.WaitGroup
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, host := range URLListIMDC1 {
			resultIMDC1, errIMDC1 = s.getAllTanks(ctx, host)
			if len(resultIMDC1) > 0 {
				for i := range resultIMDC1 {
					resultIMDC1[i].IMDC = 1
					resultIMDC1[i].MegaBoxID = int64(i/6 + 1)
				}
				break
			}
		}
	}()
	go func() {
		defer wg.Done()
		for _, host := range URLListIMDC2 {
			resultIMDC2, errIMDC2 = s.getAllTanks(ctx, host)
			if len(resultIMDC2) > 0 {
				for i := range resultIMDC2 {
					resultIMDC2[i].IMDC = 2
					resultIMDC2[i].MegaBoxID = int64(i/6 + 1)
				}
				break
			}
		}
	}()
	wg.Wait()
	if errIMDC1 != nil && len(resultIMDC1) == 0 {
		for i := 0; i < 24; i++ {
			resultIMDC1 = append(resultIMDC1, entities.Tank{
				TankID:    int64(i + 1),
				MegaBoxID: int64(i/6 + 1),
				IMDC:      1,
				Error:     errIMDC1.Error(),
			})
		}
	}
	if errIMDC2 != nil && len(resultIMDC2) == 0 {
		for i := 0; i < 24; i++ {
			resultIMDC2 = append(resultIMDC2, entities.Tank{
				TankID:    int64(i + 1),
				MegaBoxID: int64(i/6 + 1),
				IMDC:      2,
				Error:     errIMDC2.Error(),
			})
		}
	}
	return append(resultIMDC1, resultIMDC2...), nil
}

func (s *Service) getAllTanks(ctx context.Context, baseURL string) ([]entities.Tank, error) {
	res, code, err := requests.GetCurl[[]entities.Tank](ctx, fmt.Sprintf("%s/api/tank/getAllTanks", baseURL), map[string]string{
		"auth-token": s.getToken(baseURL),
	}, requests.WithSkipTLSVerification(true))
	if err != nil {
		return nil, fmt.Errorf("failed to get all tanks: %w", err)
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("wrong code in response of megaboxes: %d", code)
	}
	return res, nil
}

func (s *Service) GetVFDFan(ctx context.Context, tankID int64, baseURL string) (*entities.FanVFD, error) {
	res, code, err := requests.GetCurl[entities.FanVFD](ctx, fmt.Sprintf("%s/api/vfd/fanVFD/%d", baseURL, tankID), map[string]string{
		"auth-token": s.getToken(baseURL),
	}, requests.WithSkipTLSVerification(true))
	if err != nil {
		return nil, fmt.Errorf("failed to get VFDFan: %w", err)
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("wrong code in response of VFDFan: %d", code)
	}
	return &res, nil
}

func (s *Service) GetVFDPump(ctx context.Context, tankID int64, baseURL string) (*entities.PumpVFD, error) {
	res, code, err := requests.GetCurl[entities.PumpVFD](ctx, fmt.Sprintf("%s/api/vfd/pumpVFD/%d", baseURL, tankID), map[string]string{
		"auth-token": s.getToken(baseURL),
	}, requests.WithSkipTLSVerification(true))
	if err != nil {
		return nil, fmt.Errorf("failed to get VFDPump: %w", err)
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("wrong code in response of VFDPump: %d", code)
	}
	return &res, nil
}
