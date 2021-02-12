package config

import (
	"fmt"
	"os"
	"time"
)

func ParseTimeDuration(sTime string) time.Duration {
	timeDuration, err := time.ParseDuration(sTime)
	if err != nil {
		fmt.Println("Error while parsing time duration")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	return timeDuration
}
