package config

type Config struct {
	MetricsEndpoint string `mapstructure:"metricsEndpoint"`
	ScrapeInterval  int    `mapstructure:"scrapeInterval"`
	Port            int    `mapstructure:"port"`
	Development     bool   `mapstructure:"development"`
}
