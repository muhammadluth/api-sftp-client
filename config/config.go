package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"api-sftp-client/model"

	"github.com/joho/godotenv"
)

func LoadConfig() model.Properties {
	timestart := time.Now()
	fmt.Println("Starting Load Config " + timestart.Format("2006-01-02 15:04:05"))
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err, "init get config")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	parsePoolSize, err := strconv.Atoi(os.Getenv("POOL_CONNECTION"))
	if err != nil {
		fmt.Println(err, "init get config")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	properties := model.Properties{
		Port:     os.Getenv("PORT"),
		Timeout:  os.Getenv("TIMEOUT"),
		LogPath:  os.Getenv("LOG_PATH"),
		PoolSize: parsePoolSize,
		SftpConfig: model.SftpConfig{
			Host:     os.Getenv("SFTP_HOST"),
			User:     os.Getenv("SFTP_USER"),
			Password: os.Getenv("SFTP_PASSWORD"),
		},
	}
	timefinish := time.Now()
	fmt.Println("Finish Load Config " + timefinish.Format("2006-01-02 15:04:05"))
	return properties
}
