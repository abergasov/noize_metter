package utils

import (
	"strconv"
	"time"
)

func TimeToDayInt(src time.Time) string {
	return src.Format("20060102")
}

func TimeToDayIntNum(src time.Time) int64 {
	i, _ := strconv.ParseInt(TimeToDayInt(src), 10, 64) // we can skip err here
	return i
}
