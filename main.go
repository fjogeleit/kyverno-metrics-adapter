package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fjogeleit/kyverno-metrics-adapter/cmd"
)

func main() {
	go fakeMetricServer()

	if err := cmd.NewCLI().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func fakeMetricServer() error {
	response := `# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
	# TYPE go_memstats_stack_inuse_bytes gauge
	go_memstats_stack_inuse_bytes 425984
	# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
	# TYPE go_memstats_stack_sys_bytes gauge
	go_memstats_stack_sys_bytes 425984
	# HELP go_memstats_sys_bytes Number of bytes obtained from system.
	# TYPE go_memstats_sys_bytes gauge
	go_memstats_sys_bytes 7.36146e+07
	# HELP go_threads Number of OS threads created.
	# TYPE go_threads gauge
	go_threads 7
	# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
	# TYPE promhttp_metric_handler_requests_in_flight gauge
	promhttp_metric_handler_requests_in_flight 1
	# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
	# TYPE promhttp_metric_handler_requests_total counter
	promhttp_metric_handler_requests_total{code="200"} 0
	promhttp_metric_handler_requests_total{code="500"} 0
	promhttp_metric_handler_requests_total{code="503"} 0
	# HELP kyverno_policy_rule_results Description.
	# TYPE kyverno_policy_rule_results gauge
	kyverno_policy_rule_results{policy_validation_mode="enforce", policy_type="cluster", policy_background_mode="false", policy_name="policy-abc", policy_namespace="", resource_name="app", resource_kind="Pod", resource_namespace="default", resource_request_operation="create", rule_name="rule-abc", rule_result="FAIL", rule_type="validate", rule_execution_cause="admission_request", rule_response="message", main_request_trigger_timestamp="1619165880", rule_execution_timestamp="1619165890", policy_execution_timestamp="1619165900"} 1
	# HELP kyverno_policy_rule_count Description.
	# TYPE kyverno_policy_rule_count gauge
	kyverno_policy_rule_count{policy_validation_mode="enforce", policy_type="cluster", policy_background_mode="false", policy_name="policy-abc", policy_namespace="", rule_name="rule-abc"} 1
	kyverno_policy_rule_count{policy_validation_mode="enforce", policy_type="cluster", policy_background_mode="false", policy_name="policy-def", policy_namespace="", rule_name="rule-def"} 0
	# HELP kyverno_policy_changes Description.
	# TYPE kyverno_policy_changes gauge
	kyverno_policy_changes{policy_validation_mode="enforce", policy_type="cluster", policy_background_mode="false", policy_name="policy-abc", policy_namespace="", policy_change_type="create"} 1
	# HELP kyverno_policy_rule_execution_latency Description.
	# TYPE kyverno_policy_rule_execution_latency gauge
	kyverno_policy_rule_execution_latency{policy_validation_mode="enforce", policy_type="cluster", policy_background_mode="false", policy_name="policy-abc", policy_namespace="", resource_name="app", resource_kind="Pod", resource_namespace="default", resource_request_operation="create", rule_name="rule-abc", rule_result="FAIL", rule_type="validate", rule_execution_cause="admission_request", rule_response="message"} 16191650
	# HELP kyverno_admission_review_latency Description.
	# TYPE kyverno_admission_review_latency gauge
	kyverno_admission_review_latency{cluster_policies_count="2", namespaced_policies_count="0", validate_rules_count="2", mutate_rules_count="0", generate_rules_count="0", resource_name="app", resource_kind="Pod", resource_namespace="default", resource_request_operation="create"} 16193650
	`

	server := http.NewServeMux()
	server.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(response))
		w.Header().Add("Content-Type", "text/plain")
	})

	err := http.ListenAndServe(":2112", server)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
