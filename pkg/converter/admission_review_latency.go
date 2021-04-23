package converter

import (
	"strconv"

	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/kyverno"
	prometheus "github.com/prometheus/client_model/go"
)

func convertAdmissionReviewLatency(rawMetrics *prometheus.MetricFamily) []kyverno.AdmissionReviewLatency {
	list := make([]kyverno.AdmissionReviewLatency, 0, len(rawMetrics.Metric))

	for _, metric := range rawMetrics.Metric {
		values := convertLabels(metric.Label)
		resource := convertResource(values)

		arl := kyverno.AdmissionReviewLatency{
			Resource: resource,
		}

		arl.Latency = int(*metric.Gauge.Value)

		if value, ok := values["cluster_policies_count"]; ok {
			count, err := strconv.Atoi(value)
			if err == nil {
				arl.ClusterPoliciesCount = count
			}
		}

		if value, ok := values["namespaced_policies_count"]; ok {
			count, err := strconv.Atoi(value)
			if err == nil {
				arl.NamespacedPoliciesCount = count
			}
		}

		if value, ok := values["validate_rules_count"]; ok {
			count, err := strconv.Atoi(value)
			if err == nil {
				arl.ValidateRulesCount = count
			}
		}

		if value, ok := values["mutate_rules_count"]; ok {
			count, err := strconv.Atoi(value)
			if err == nil {
				arl.MutateRulesCount = count
			}
		}

		if value, ok := values["generate_rules_count"]; ok {
			count, err := strconv.Atoi(value)
			if err == nil {
				arl.GenerateRulesCount = count
			}
		}

		list = append(list, arl)
	}

	return list
}
