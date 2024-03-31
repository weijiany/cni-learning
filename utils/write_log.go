package utils

import (
	"os"
)

const (
	LogPath = "./info.txt"
)

func WriteLog(contents ...string) {
	for _, content := range contents {
		file, _ := os.OpenFile(LogPath, os.O_APPEND|os.O_WRONLY, 0644)
		defer file.Close()

		file.WriteString(content + "\n")
	}
}

func init() {
	os.Truncate(LogPath, 0)
}
