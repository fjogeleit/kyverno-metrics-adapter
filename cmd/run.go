package cmd

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/api"
	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/client"
	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/converter"
	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/kyverno"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func newRunCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Scrap Metrics and start REST API",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := loadConfig(cmd)
			if err != nil {
				return err
			}

			if c.Development {
				fmt.Println("Start in development mode")
			}

			fmt.Printf("ScrapeInterval: %ds\n", c.ScrapeInterval)

			store := kyverno.NewStore()
			client := client.NewClient(c.MetricsEndpoint)

			loadMetrics(store, client)

			g := new(errgroup.Group)

			g.Go(func() error {
				for {
					time.Sleep(time.Duration(c.ScrapeInterval) * time.Second)

					err := loadMetrics(store, client)
					if err != nil {
						return err
					}
					log.Println("Scrape Metric successful")
				}
			})

			server := api.NewServer(store, c.Port, c.Development)

			g.Go(server.Start)

			return g.Wait()
		},
	}

	cmd.PersistentFlags().StringP("kyverno-metrics-endpoint", "m", "http://kyverno.kyverno.svc.cluster.local:2112/metrics", "HTTP Endpoint for Kyverno Metrics")
	cmd.PersistentFlags().IntP("port", "p", 3000, "REST API Port")
	cmd.PersistentFlags().BoolP("development", "d", false, "Enable CORS Header")
	cmd.PersistentFlags().IntP("scrape-interval", "s", 5, "Metrics Scrape Interval in Seconds")

	flag.Parse()

	return cmd
}

func loadMetrics(store kyverno.Store, client client.Client) error {
	body, err := client.FetchMetrics()
	if err != nil {
		return err
	}

	metrics, err := converter.Convert(body)
	if err != nil {
		return err
	}

	store.SetMetrics(metrics)

	return nil
}
