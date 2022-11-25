package utils

import (
	"fmt"
	"time"
)

func ConvertTimeStampToUnixMicro(timestamp string) int64 {
	layout := "2006-01-02T15:04:05.9-08:00"
	t, err := time.Parse(layout, timestamp)
	if err != nil {
		fmt.Println(err)
	}
	return t.UnixMicro()
}
