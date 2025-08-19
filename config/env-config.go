package config

import "github.com/spf13/viper"

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
}

func SetupEnvConfig() (*EnvConfig, error) {
	viper.SetConfigName("env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var envConf EnvConfig
	if err := viper.Unmarshal(&envConf); err != nil {
		return nil, err
	}

	return &envConf, nil
}
