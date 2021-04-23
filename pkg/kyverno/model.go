package kyverno

type PolicyViolationMode string

type PolicyType string

type ResourceRequestOperation string

type RuleResult string

type RuleType string

type RuleExecutionCause string

type PolicyChangeType string

const (
	Audit   PolicyViolationMode = "audit"
	Enforce PolicyViolationMode = "enforce"

	Cluster    PolicyType = "cluster"
	Namespaced PolicyType = "namespaced"

	Create ResourceRequestOperation = "create"
	Update ResourceRequestOperation = "update"
	Delete ResourceRequestOperation = "delete"

	Pass    RuleResult = "PASS"
	Fail    RuleResult = "FAIL"
	Error   RuleResult = "ERROR"
	Skip    RuleResult = "SKIP"
	Warning RuleResult = "WARN"

	Validate RuleType = "validate"
	Mutate   RuleType = "mutate"
	Generate RuleType = "generate"

	AdmissionRequest RuleExecutionCause = "admission_request"
	BackgroundScan   RuleExecutionCause = "background_scan"

	PolicyCreate PolicyChangeType = "create"
	PolicyUpdate PolicyChangeType = "update"
	PolicyDelete PolicyChangeType = "delete"
)

type Policy struct {
	Name           string              `json:"name"`
	Namespace      string              `json:"namespace,omitempty"`
	Type           PolicyType          `json:"type"`
	ValidationMode PolicyViolationMode `json:"validationMode"`
	BackgroundMode bool                `json:"backgroundMode"`
}

type Resource struct {
	Name             string                   `json:"name"`
	Kind             string                   `json:"kind"`
	Namespace        string                   `json:"namespace,omitempty"`
	RequestOperation ResourceRequestOperation `json:"requestOperation"`
}

type Rule struct {
	Name           string             `json:"name"`
	Response       string             `json:"response,omitempty"`
	Result         RuleResult         `json:"result"`
	Type           RuleType           `json:"type"`
	ExecutionCause RuleExecutionCause `json:"executionCause"`
}

type PolicyRuleResult struct {
	Policy                      Policy   `json:"policy"`
	Resource                    Resource `json:"resource"`
	Rule                        Rule     `json:"rule"`
	MainRequestTriggerTimestamp int      `json:"mainRequestTriggerTimestamp"`
	RuleExecutionTimestamp      int      `json:"ruleExecutionTimestamp"`
	PolicyExecutionTimestamp    int      `json:"policyExecutionTimestamp"`
}

type PolicyRuleCount struct {
	Policy   Policy `json:"policy"`
	RuleName string `json:"ruleName"`
}

type PolicyChange struct {
	Policy           Policy           `json:"policy"`
	PolicyChangeType PolicyChangeType `json:"policyChangeType"`
}

type PolicyRuleExecutionLatency struct {
	Policy   Policy   `json:"policy"`
	Resource Resource `json:"resource"`
	Rule     Rule     `json:"rule"`
	Latency  int      `json:"latency"`
}

type AdmissionReviewLatency struct {
	ClusterPoliciesCount    int      `json:"clusterPoliciesCount"`
	NamespacedPoliciesCount int      `json:"namespacedPoliciesCount"`
	ValidateRulesCount      int      `json:"validateRulesCount"`
	MutateRulesCount        int      `json:"mutateRulesCount"`
	GenerateRulesCount      int      `json:"generateRulesCount"`
	Resource                Resource `json:"resource"`
	Latency                 int      `json:"latency"`
}

type Metrics struct {
	PolicyRuleResults          []PolicyRuleResult
	PolicyRuleCount            []PolicyRuleCount
	PolicyChanges              []PolicyChange
	PolicyRuleExecutionLatency []PolicyRuleExecutionLatency
	AdmissionReviewLatency     []AdmissionReviewLatency
}
