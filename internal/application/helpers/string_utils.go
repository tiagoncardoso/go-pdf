package helpers

import (
	"strconv"
	"time"
)

func GenerateFileName(prefix string) string {
	timestamp := time.Now().Unix()
	return prefix + "_" + strconv.FormatInt(timestamp, 10) + ".pdf"
}
