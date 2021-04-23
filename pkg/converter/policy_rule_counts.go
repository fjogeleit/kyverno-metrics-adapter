package converter

import (
	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/kyverno"
	prometheus "github.com/prometheus/client_model/go"
)

func convertPolicyRuleCount(rawMetrics *prometheus.MetricFamily) []kyverno.PolicyRuleCount {
	counts := make([]kyverno.PolicyRuleCount, 0)

	for _, metric := range rawMetrics.Metric {
		if *metric.Gauge.Value == 0 {
			continue
		}

		values := convertLabels(metric.Label)
		policy := convertPolicy(values)
		rule := convertRule(values)

		count := kyverno.PolicyRuleCount{
			Policy:   policy,
			RuleName: rule.Name,
		}

		counts = append(counts, count)
	}

	return counts
}
