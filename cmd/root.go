package cmd

import (
	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCLI creates a new instance of the root CLI
func NewCLI() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "kyverno-metrics-adapter",
		Short: "Converts Kyverno Metrics into Restful APIs",
		Long:  `Converts Kyverno Metrics into Restful APIs`,
	}

	rootCmd.AddCommand(newRunCMD())

	return rootCmd
}

func loadConfig(cmd *cobra.Command) (*config.Config, error) {
	v := viper.New()

	v.AutomaticEnv()

	if flag := cmd.Flags().Lookup("scrape-interval"); flag != nil {
		v.BindPFlag("scrapeInterval", flag)
	}
	if flag := cmd.Flags().Lookup("kyverno-metrics-endpoint"); flag != nil {
		v.BindPFlag("metricsEndpoint", flag)
	}
	if flag := cmd.Flags().Lookup("port"); flag != nil {
		v.BindPFlag("port", flag)
	}
	if flag := cmd.Flags().Lookup("development"); flag != nil {
		v.BindPFlag("development", flag)
	}

	c := &config.Config{}

	err := v.Unmarshal(c)

	return c, err
}
