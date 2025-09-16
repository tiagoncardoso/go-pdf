package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/tiagoncardoso/go-pdf/pkg/logger"
)

type EnvConfig struct {
	AppPort                  string `mapstructure:"APP_PORT"`
	Dpi                      uint   `mapstructure:"DPI"`
	OutputPath               string `mapstructure:"OUTPUT_PATH"`
	Title                    string `mapstructure:"TITLE"`
	Orientation              string `mapstructure:"ORIENTATION"`
	PageSize                 string `mapstructure:"PAGE_SIZE"`
	StorageEndpoint          string `mapstructure:"STORAGE_ENDPOINT"`
	StorageSpaceName         string `mapstructure:"STORAGE_SPACE_NAME"`
	StorageAccessKey         string `mapstructure:"STORAGE_ACCESS_KEY"`
	StorageSecretKey         string `mapstructure:"STORAGE_SECRET_KEY"`
	StorageRegion            string `mapstructure:"STORAGE_REGION"`
	ReportPrefix             string `mapstructure:"REPORT_PREFIX"`
	PdfLinkExpirationSeconds int    `mapstructure:"PDF_LINK_EXPIRATION_SECONDS"`
	BasicAuthRealm           string `mapstructure:"BASIC_AUTH_REALM"`
	BasicAuthClientID        string `mapstructure:"BASIC_AUTH_CLIENT_ID"`
	BasicAuthClientSecret    string `mapstructure:"BASIC_AUTH_CLIENT_SECRET"`
}

func SetupEnvConfig() (*EnvConfig, error) {
	// Load .env file if present, ignore error if not
	_ = godotenv.Load()

	cfg := &EnvConfig{
		AppPort:                  os.Getenv("APP_PORT"),
		Dpi:                      parseUintEnv("DPI"),
		OutputPath:               os.Getenv("OUTPUT_PATH"),
		Title:                    os.Getenv("TITLE"),
		Orientation:              os.Getenv("ORIENTATION"),
		PageSize:                 os.Getenv("PAGE_SIZE"),
		StorageEndpoint:          os.Getenv("STORAGE_ENDPOINT"),
		StorageSpaceName:         os.Getenv("STORAGE_SPACE_NAME"),
		StorageAccessKey:         os.Getenv("STORAGE_ACCESS_KEY"),
		StorageSecretKey:         os.Getenv("STORAGE_SECRET_KEY"),
		StorageRegion:            os.Getenv("STORAGE_REGION"),
		ReportPrefix:             os.Getenv("REPORT_PREFIX"),
		PdfLinkExpirationSeconds: parseIntEnv("PDF_LINK_EXPIRATION_SECONDS"),
		BasicAuthRealm:           os.Getenv("BASIC_AUTH_REALM"),
		BasicAuthClientID:        os.Getenv("BASIC_AUTH_CLIENT_ID"),
		BasicAuthClientSecret:    os.Getenv("BASIC_AUTH_CLIENT_SECRET"),
	}

	return cfg, nil
}

func parseUintEnv(key string) uint {
	val := os.Getenv(key)
	if val == "" {
		return 0
	}
	var u uint64
	u, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		logger.Warn("Invalid uint value for %s: %v", key, err)
		return 0
	}
	return uint(u)
}

func parseIntEnv(key string) int {
	val := os.Getenv(key)
	if val == "" {
		return 0
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		logger.Warn("Invalid int value for %s: %v", key, err)
		return 0
	}
	return i
}
