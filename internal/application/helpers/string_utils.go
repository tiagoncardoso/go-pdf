package helpers

import (
	"strconv"
	"time"
)

func GenerateFileName(prefix string) string {
	// Generate a unique file name using the current timestamp
	timestamp := time.Now().Unix()
	return prefix + "_" + strconv.FormatInt(timestamp, 10) + ".pdf"
}
