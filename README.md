# Kyverno Metrics Adapter

Converts Metrics provided by Kyverno into restful APIs

```bash
Scrap Metrics and start REST API

Usage:
  kyverno-metrics-adapter run [flags]

Flags:
  -d, --development                       Enable CORS Header
  -h, --help                              help for run
  -m, --kyverno-metrics-endpoint string   HTTP Endpoint for Kyverno Metrics (default "http://kyverno.kyverno.svc.cluster.local:2112/metrics")
  -p, --port int                          REST API Port (default 3000)
  -s, --scrape-interval int               Metrics Scrape Interval in Seconds (default 5)
```

## Policy Rule Results

### GET API Response

```bash
http://localhost:3000/policy-rule-results
```

### Example Metric

```txt
kyverno_policy_rule_results{policy_validation_mode="enforce", policy_type="cluster", policy_background_mode="false", policy_name="policy-abc", policy_namespace="", resource_name="app", resource_kind="Pod", resource_namespace="default", resource_request_operation="create", rule_name="rule-abc", rule_result="FAIL", rule_type="validate", rule_execution_cause="admission_request", rule_response="message", main_request_trigger_timestamp="1619165880", rule_execution_timestamp="1619165890", policy_execution_timestamp="1619165900"} 1
```

### GET API Response

```json
[
  {
    "policy": {
      "name": "policy-abc",
      "type": "cluster",
      "validationMode": "enforce",
      "backgroundMode": false
    },
    "resource": {
      "name": "app",
      "kind": "Pod",
      "namespace": "default",
      "requestOperation": "create"
    },
    "rule": {
      "name": "rule-abc",
      "response": "message",
      "result": "FAIL",
      "type": "validate",
      "executionCause": "admission_request"
    },
    "mainRequestTriggerTimestamp": 1619165880,
    "ruleExecutionTimestamp": 1619165890,
    "policyExecutionTimestamp": 1619165900
  }
]
```

## Policy Rule Count

Filters items with `0` value.

### GET API Response

```bash
http://localhost:3000/policy-rule-count
```

### Example Metric

```txt
kyverno_policy_rule_count{policy_validation_mode="enforce", policy_type="cluster", policy_background_mode="false", policy_name="policy-abc", policy_namespace="", rule_name="rule-abc"} 1
kyverno_policy_rule_count{policy_validation_mode="enforce", policy_type="cluster", policy_background_mode="false", policy_name="policy-def", policy_namespace="", rule_name="rule-def"} 0
```

### GET API Response

```json
[
  {
    "policy": {
      "name": "policy-abc",
      "type": "cluster",
      "validationMode": "enforce",
      "backgroundMode": false
    },
    "ruleName": "rule-abc"
  }
]
```

## Policy Changes

### GET API Response

```bash
http://localhost:3000/policy-changes
```

### Example Metric

```txt
kyverno_policy_changes{policy_validation_mode="enforce", policy_type="cluster", policy_background_mode="false", policy_name="policy-abc", policy_namespace="", policy_change_type="create"} 1
```

### GET API Response

```json
[
  {
    "policy": {
      "name": "policy-abc",
      "type": "cluster",
      "validationMode": "enforce",
      "backgroundMode": false
    },
    "policyChangeType": "create"
  }
]
```

## Policy Rule Exeution Latency

### GET API

```bash
http://localhost:3000/policy-rule-execution-latency
```

### Example Metric

```txt
kyverno_policy_rule_execution_latency{policy_validation_mode="enforce", policy_type="cluster", policy_background_mode="false", policy_name="policy-abc", policy_namespace="", resource_name="app", resource_kind="Pod", resource_namespace="default", resource_request_operation="create", rule_name="rule-abc", rule_result="FAIL", rule_type="validate", rule_execution_cause="admission_request", rule_response="message"} 16191650
```

### GET API Response

```json
[
  {
    "policy": {
      "name": "policy-abc",
      "type": "cluster",
      "validationMode": "enforce",
      "backgroundMode": false
    },
    "resource": {
      "name": "app",
      "kind": "Pod",
      "namespace": "default",
      "requestOperation": "create"
    },
    "rule": {
      "name": "rule-abc",
      "response": "message",
      "result": "FAIL",
      "type": "validate",
      "executionCause": "admission_request"
    },
    "latency": 16191650
  }
]
```

## Admission Review Latency

### GET API

```bash
http://localhost:3000/admission-review-latency
```

### Example Metric

```txt
kyverno_admission_review_latency{cluster_policies_count="2", namespaced_policies_count="0", validate_rules_count="2", mutate_rules_count="0", generate_rules_count="0", resource_name="app", resource_kind="Pod", resource_namespace="default", resource_request_operation="create"} 16193650
```

### GET API Response

```json
[
  {
    "clusterPoliciesCount": 2,
    "namespacedPoliciesCount": 0,
    "validateRulesCount": 2,
    "mutateRulesCount": 0,
    "generateRulesCount": 0,
    "resource": {
      "name": "app",
      "kind": "Pod",
      "namespace": "default",
      "requestOperation": "create"
    },
    "latency": 16193650
  }
]
```
