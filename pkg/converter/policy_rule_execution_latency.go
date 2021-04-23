package converter

import (
	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/kyverno"
	prometheus "github.com/prometheus/client_model/go"
)

func convertPolicyRuleExecutionLatency(rawMetrics *prometheus.MetricFamily) []kyverno.PolicyRuleExecutionLatency {
	results := make([]kyverno.PolicyRuleExecutionLatency, 0, len(rawMetrics.Metric))

	for _, metric := range rawMetrics.Metric {
		values := convertLabels(metric.Label)
		policy := convertPolicy(values)
		resource := convertResource(values)
		rule := convertRule(values)

		result := kyverno.PolicyRuleExecutionLatency{
			Policy:   policy,
			Resource: resource,
			Rule:     rule,
		}

		result.Latency = int(*metric.Gauge.Value)

		results = append(results, result)
	}

	return results
}
