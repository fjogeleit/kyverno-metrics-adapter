package converter

import (
	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/kyverno"
	prometheus "github.com/prometheus/client_model/go"
)

func convertPolicyChanges(rawMetrics *prometheus.MetricFamily) []kyverno.PolicyChange {
	changes := make([]kyverno.PolicyChange, 0)

	for _, metric := range rawMetrics.Metric {
		if *metric.Gauge.Value == 0 {
			continue
		}

		values := convertLabels(metric.Label)
		policy := convertPolicy(values)

		change := kyverno.PolicyChange{
			Policy: policy,
		}

		if entries, ok := values["policy_change_type"]; ok {
			change.PolicyChangeType = kyverno.PolicyChangeType(entries)
		}

		changes = append(changes, change)
	}

	return changes
}
