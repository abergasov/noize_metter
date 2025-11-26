package ces

import (
	"context"
	"fmt"
	"noize_metter/internal/entities"
	"noize_metter/internal/utils"
	"time"
)

func (s *Service) RunIteration() error {
	s.tokenMU.RLock()
	l := len(s.token)
	s.tokenMU.RUnlock()
	if l == 0 {
		s.log.Info("token is empty, filling token")
		if err := s.FillToken(); err != nil {
			return fmt.Errorf("failed to create token: %w", err)
		}
	}
	ctx, cancel := context.WithTimeout(s.ctx, 10*time.Minute)
	defer cancel()
	megaBoxes, err := s.GetAllMegaBoxes(ctx)
	if err != nil {
		return fmt.Errorf("failed to get all mega boxes: %w", err)
	}
	dt := utils.RoundToNearest5Minutes(time.Now())
	for i := range megaBoxes {
		megaBoxes[i].TimestampNum = utils.TimeToDayIntNum(dt)
		megaBoxes[i].TimestampPQ = dt.Format(time.DateTime)
	}

	if err = s.repo.SaveMegaBoxes(megaBoxes); err != nil {
		return fmt.Errorf("failed to save mega boxes: %w", err)
	}

	tanks, err := s.GetAllTanks(ctx)
	if err != nil {
		return fmt.Errorf("failed to get all tanks: %w", err)
	}

	convertedTanks := entities.ConvertAPITanks(dt, tanks)
	if err = s.repo.SaveTanks(convertedTanks); err != nil {
		return fmt.Errorf("failed to save tanks: %w", err)
	}

	cesChannelsTables := make([]entities.CesTanksChannels, 0, len(tanks)*8*6)
	cesChannelsTablesV2 := make([]entities.CesTanksChannelsV2, 0, len(tanks)*8*6*8)
	for i := range tanks {
		if len(tanks[i].SPDU.Sections) < 8 {
			continue
		}
		for j := range tanks[i].SPDU.Sections {
			if len(tanks[i].SPDU.Sections[j].Channels) < 6 {
				continue
			}
		}

		for j := 0; j <= 7; j++ {
			for k := 0; k <= 5; k++ {
				cesChannelsTablesV2 = append(cesChannelsTablesV2, entities.CesTanksChannelsV2{
					TimestampNum: utils.TimeToDayIntNum(dt),
					TimestampPQ:  dt.Format(time.DateTime),
					TankID:       tanks[i].TankID,
					IMDC:         tanks[i].IMDC,
					MegaboxID:    tanks[i].MegaBoxID,
					Section:      int64(j + 1),
					Channel:      int64(k + 1),
					Amperage:     tanks[i].SPDU.Sections[j].Channels[k].Rms,
				})
			}
		}

		tmp := entities.CesTanksChannels{
			TimestampNum: utils.TimeToDayIntNum(dt),
			TimestampPQ:  dt.Format(time.DateTime),
			TankID:       tanks[i].TankID,
			IMDC:         tanks[i].IMDC,
			MegaboxID:    tanks[i].MegaBoxID,
			// Section 1 Channels
			Section1Channel1: tanks[i].SPDU.Sections[0].Channels[0].Rms,
			Section1Channel2: tanks[i].SPDU.Sections[0].Channels[1].Rms,
			Section1Channel3: tanks[i].SPDU.Sections[0].Channels[2].Rms,
			Section1Channel4: tanks[i].SPDU.Sections[0].Channels[3].Rms,
			Section1Channel5: tanks[i].SPDU.Sections[0].Channels[4].Rms,
			Section1Channel6: tanks[i].SPDU.Sections[0].Channels[5].Rms,
			// Section 2 Channels
			Section2Channel1: tanks[i].SPDU.Sections[1].Channels[0].Rms,
			Section2Channel2: tanks[i].SPDU.Sections[1].Channels[1].Rms,
			Section2Channel3: tanks[i].SPDU.Sections[1].Channels[2].Rms,
			Section2Channel4: tanks[i].SPDU.Sections[1].Channels[3].Rms,
			Section2Channel5: tanks[i].SPDU.Sections[1].Channels[4].Rms,
			Section2Channel6: tanks[i].SPDU.Sections[1].Channels[5].Rms,
			// Section 3 Channels
			Section3Channel1: tanks[i].SPDU.Sections[2].Channels[0].Rms,
			Section3Channel2: tanks[i].SPDU.Sections[2].Channels[1].Rms,
			Section3Channel3: tanks[i].SPDU.Sections[2].Channels[2].Rms,
			Section3Channel4: tanks[i].SPDU.Sections[2].Channels[3].Rms,
			Section3Channel5: tanks[i].SPDU.Sections[2].Channels[4].Rms,
			Section3Channel6: tanks[i].SPDU.Sections[2].Channels[5].Rms,
			// Section 4 Channels
			Section4Channel1: tanks[i].SPDU.Sections[3].Channels[0].Rms,
			Section4Channel2: tanks[i].SPDU.Sections[3].Channels[1].Rms,
			Section4Channel3: tanks[i].SPDU.Sections[3].Channels[2].Rms,
			Section4Channel4: tanks[i].SPDU.Sections[3].Channels[3].Rms,
			Section4Channel5: tanks[i].SPDU.Sections[3].Channels[4].Rms,
			Section4Channel6: tanks[i].SPDU.Sections[3].Channels[5].Rms,
			// Section 5 Channels
			Section5Channel1: tanks[i].SPDU.Sections[4].Channels[0].Rms,
			Section5Channel2: tanks[i].SPDU.Sections[4].Channels[1].Rms,
			Section5Channel3: tanks[i].SPDU.Sections[4].Channels[2].Rms,
			Section5Channel4: tanks[i].SPDU.Sections[4].Channels[3].Rms,
			Section5Channel5: tanks[i].SPDU.Sections[4].Channels[4].Rms,
			Section5Channel6: tanks[i].SPDU.Sections[4].Channels[5].Rms,
			// Section 6 Channels
			Section6Channel1: tanks[i].SPDU.Sections[5].Channels[0].Rms,
			Section6Channel2: tanks[i].SPDU.Sections[5].Channels[1].Rms,
			Section6Channel3: tanks[i].SPDU.Sections[5].Channels[2].Rms,
			Section6Channel4: tanks[i].SPDU.Sections[5].Channels[3].Rms,
			Section6Channel5: tanks[i].SPDU.Sections[5].Channels[4].Rms,
			Section6Channel6: tanks[i].SPDU.Sections[5].Channels[5].Rms,
			// Section 7 Channels
			Section7Channel1: tanks[i].SPDU.Sections[6].Channels[0].Rms,
			Section7Channel2: tanks[i].SPDU.Sections[6].Channels[1].Rms,
			Section7Channel3: tanks[i].SPDU.Sections[6].Channels[2].Rms,
			Section7Channel4: tanks[i].SPDU.Sections[6].Channels[3].Rms,
			Section7Channel5: tanks[i].SPDU.Sections[6].Channels[4].Rms,
			Section7Channel6: tanks[i].SPDU.Sections[6].Channels[5].Rms,
			// Section 8 Channels
			Section8Channel1: tanks[i].SPDU.Sections[7].Channels[0].Rms,
			Section8Channel2: tanks[i].SPDU.Sections[7].Channels[1].Rms,
			Section8Channel3: tanks[i].SPDU.Sections[7].Channels[2].Rms,
			Section8Channel4: tanks[i].SPDU.Sections[7].Channels[3].Rms,
			Section8Channel5: tanks[i].SPDU.Sections[7].Channels[4].Rms,
			Section8Channel6: tanks[i].SPDU.Sections[7].Channels[5].Rms,
		}
		cesChannelsTables = append(cesChannelsTables, tmp)
	}
	if err = s.repo.SaveCesTanksChannelsV2(cesChannelsTablesV2); err != nil {
		return fmt.Errorf("failed to save ces tanks channels v2: %w", err)
	}
	if err = s.repo.SaveCesTanksChannels(cesChannelsTables); err != nil {
		return fmt.Errorf("failed to save ces tanks channels: %w", err)
	}
	return nil
}
