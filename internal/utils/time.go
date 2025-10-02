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
func RoundToNearest5Minutes(t time.Time) time.Time {
	roundedMinutes := (t.Minute() / 5) * 5
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), roundedMinutes, 0, 0, t.Location())
}
