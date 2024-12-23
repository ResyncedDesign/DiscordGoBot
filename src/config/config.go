package config

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

var startTime time.Time

func Load() {
	startTime = time.Now()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func Uptime() time.Duration {
	return time.Since(startTime)
}

func FormattedUptime() string {
	duration := time.Since(startTime)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02dh:%02dm:%02ds", hours, minutes, seconds)
}
