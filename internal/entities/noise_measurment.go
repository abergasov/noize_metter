package entities

import "time"

type NoiseMeasures struct {
	Timestamp        time.Time `db:"timestamp" json:"timestamp"`
	TimestampPQ      string    `json:"-" parquet:"name=timestamp, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	TimestampNum     int64     `json:"timestamp_num" db:"timestamp_num" parquet:"name=timestamp_num, type=INT64"`
	TimestampHourNum int64     `db:"timestamp_hour_num" json:"-" parquet:"name=timestamp_hour_num, type=INT64"`

	LAeqDT  float64 `db:"laeq_dt" json:"laeq_dt" parquet:"name=laeq_dt, type=DOUBLE"`
	LAf     float64 `db:"laf" json:"laf" parquet:"name=laf, type=DOUBLE"`
	LCPK    float64 `db:"lcpk" json:"lcpk" parquet:"name=lcpk, type=DOUBLE"`
	LAeqG10 float64 `db:"laeq_g10" json:"laeq_g10" parquet:"name=laeq_g10, type=DOUBLE"`
	LAeqG5  float64 `db:"laeq_g5" json:"laeq_g5" parquet:"name=laeq_g5, type=DOUBLE"`
}
