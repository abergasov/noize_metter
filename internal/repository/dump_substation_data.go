package repository

import "noize_metter/internal/entities"

func (r *Repo) DumpSubstationRawData(items []entities.ModbusRegisters) error {
	return saveItems(r.conf.StorageSubstationFolder, "substation_raw", items)
}
