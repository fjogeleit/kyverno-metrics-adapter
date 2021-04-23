package converter

import (
	"io"

	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/kyverno"
	prometheus "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

func Convert(response io.Reader) (kyverno.Metrics, error) {
	metrics := kyverno.Metrics{}

	var parser expfmt.TextParser
	rawMetrics, err := parser.TextToMetricFamilies(response)
	if err != nil {
		return metrics, err
	}

	if entries, ok := rawMetrics["kyverno_policy_rule_results"]; ok {
		metrics.PolicyRuleResults = convertPolicyRuleResults(entries)
	}

	if entries, ok := rawMetrics["kyverno_policy_rule_count"]; ok {
		metrics.PolicyRuleCount = convertPolicyRuleCount(entries)
	}

	if entries, ok := rawMetrics["kyverno_policy_changes"]; ok {
		metrics.PolicyChanges = convertPolicyChanges(entries)
	}

	if entries, ok := rawMetrics["kyverno_policy_rule_execution_latency"]; ok {
		metrics.PolicyRuleExecutionLatency = convertPolicyRuleExecutionLatency(entries)
	}

	if entries, ok := rawMetrics["kyverno_admission_review_latency"]; ok {
		metrics.AdmissionReviewLatency = convertAdmissionReviewLatency(entries)
	}

	return metrics, nil
}

func convertLabels(labels []*prometheus.LabelPair) map[string]string {
	labelMap := make(map[string]string, 0)

	for _, label := range labels {
		labelMap[*label.Name] = *label.Value
	}

	return labelMap
}

func convertPolicy(values map[string]string) kyverno.Policy {
	policy := kyverno.Policy{}

	if mode, ok := values["policy_validation_mode"]; ok {
		policy.ValidationMode = kyverno.PolicyViolationMode(mode)
	}

	if policyType, ok := values["policy_type"]; ok {
		policy.Type = kyverno.PolicyType(policyType)
	}

	if background, ok := values["policy_background_mode"]; ok {
		policy.BackgroundMode = background == "true"
	}

	if namespace, ok := values["policy_namespace"]; ok {
		policy.Namespace = namespace
	}

	if name, ok := values["policy_name"]; ok {
		policy.Name = name
	}

	return policy
}

func convertResource(values map[string]string) kyverno.Resource {
	resource := kyverno.Resource{}

	if operation, ok := values["resource_request_operation"]; ok {
		resource.RequestOperation = kyverno.ResourceRequestOperation(operation)
	}

	if kind, ok := values["resource_kind"]; ok {
		resource.Kind = kind
	}

	if namespace, ok := values["resource_namespace"]; ok {
		resource.Namespace = namespace
	}

	if name, ok := values["resource_name"]; ok {
		resource.Name = name
	}

	return resource
}

func convertRule(values map[string]string) kyverno.Rule {
	rule := kyverno.Rule{}

	if name, ok := values["rule_name"]; ok {
		rule.Name = name
	}

	if response, ok := values["rule_response"]; ok {
		rule.Response = response
	}

	if result, ok := values["rule_result"]; ok {
		rule.Result = kyverno.RuleResult(result)
	}

	if ruleType, ok := values["rule_type"]; ok {
		rule.Type = kyverno.RuleType(ruleType)
	}

	if executionCause, ok := values["rule_execution_cause"]; ok {
		rule.ExecutionCause = kyverno.RuleExecutionCause(executionCause)
	}

	return rule
}
