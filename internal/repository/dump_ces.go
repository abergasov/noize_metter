package repository

import (
	"noize_metter/internal/entities"
)

func (r *Repo) SaveMegaBoxes(boxes []entities.MegaBox) error {
	return saveItems(r.conf.StorageCESMegaBoxesFolder, "ces_raw_boxes", boxes)
}

func (r *Repo) SaveTanks(boxes []entities.CesTank) error {
	return saveItems(r.conf.StorageCESTanksFolder, "ces_raw_tanks", boxes)
}

func (r *Repo) SaveCesTanksChannels(items []entities.CesTanksChannels) error {
	return saveItems(r.conf.StorageCESChannelsFolder, "ces_raw_channels", items)
}

func (r *Repo) SaveCesTanksChannelsV2(items []entities.CesTanksChannelsV2) error {
	return saveItems(r.conf.StorageCESChannelsV2Folder, "ces_raw_channels_v2", items)
}
