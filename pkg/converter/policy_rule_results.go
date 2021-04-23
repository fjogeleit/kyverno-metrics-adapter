package converter

import (
	"strconv"

	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/kyverno"
	prometheus "github.com/prometheus/client_model/go"
)

func convertPolicyRuleResults(rawMetrics *prometheus.MetricFamily) []kyverno.PolicyRuleResult {
	results := make([]kyverno.PolicyRuleResult, 0, len(rawMetrics.Metric))

	for _, metric := range rawMetrics.Metric {
		values := convertLabels(metric.Label)
		policy := convertPolicy(values)
		resource := convertResource(values)
		rule := convertRule(values)

		result := kyverno.PolicyRuleResult{
			Policy:   policy,
			Resource: resource,
			Rule:     rule,
		}

		if value, ok := values["rule_execution_timestamp"]; ok {
			timestamp, err := strconv.Atoi(value)
			if err == nil {
				result.RuleExecutionTimestamp = timestamp
			}
		}

		if value, ok := values["main_request_trigger_timestamp"]; ok {
			timestamp, err := strconv.Atoi(value)
			if err == nil {
				result.MainRequestTriggerTimestamp = timestamp
			}
		}

		if value, ok := values["policy_execution_timestamp"]; ok {
			timestamp, err := strconv.Atoi(value)
			if err == nil {
				result.PolicyExecutionTimestamp = timestamp
			}
		}

		results = append(results, result)
	}

	return results
}
