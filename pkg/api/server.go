package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fjogeleit/kyverno-metrics-adapter/pkg/kyverno"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server interface {
	Start() error
}

type server struct {
	store       kyverno.Store
	mux         *http.ServeMux
	development bool
	port        int
}

func (s *server) registerHandler() {
	s.mux.HandleFunc("/policy-rule-results", Gzip(func(w http.ResponseWriter, req *http.Request) {
		if s.development {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		list := s.store.GetMetrics().PolicyRuleResults
		if len(list) == 0 {
			fmt.Fprint(w, "[]")

			return
		}

		if err := json.NewEncoder(w).Encode(list); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{ "message": "%s" }`, err.Error())
		}
	}))

	s.mux.HandleFunc("/policy-rule-count", Gzip(func(w http.ResponseWriter, req *http.Request) {
		if s.development {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		list := s.store.GetMetrics().PolicyRuleCount
		if len(list) == 0 {
			fmt.Fprint(w, "[]")

			return
		}

		if err := json.NewEncoder(w).Encode(list); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{ "message": "%s" }`, err.Error())
		}
	}))

	s.mux.HandleFunc("/policy-changes", Gzip(func(w http.ResponseWriter, req *http.Request) {
		if s.development {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		list := s.store.GetMetrics().PolicyChanges
		if len(list) == 0 {
			fmt.Fprint(w, "[]")

			return
		}

		if err := json.NewEncoder(w).Encode(list); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{ "message": "%s" }`, err.Error())
		}
	}))

	s.mux.HandleFunc("/policy-rule-execution-latency", Gzip(func(w http.ResponseWriter, req *http.Request) {
		if s.development {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		list := s.store.GetMetrics().PolicyRuleExecutionLatency
		if len(list) == 0 {
			fmt.Fprint(w, "[]")

			return
		}

		if err := json.NewEncoder(w).Encode(list); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{ "message": "%s" }`, err.Error())
		}
	}))

	s.mux.HandleFunc("/admission-review-latency", Gzip(func(w http.ResponseWriter, req *http.Request) {
		if s.development {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		list := s.store.GetMetrics().AdmissionReviewLatency
		if len(list) == 0 {
			fmt.Fprint(w, "[]")

			return
		}

		if err := json.NewEncoder(w).Encode(list); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{ "message": "%s" }`, err.Error())
		}
	}))

	s.mux.Handle("/metrics", promhttp.Handler())
}

func (s *server) Start() error {
	s.registerHandler()

	fmt.Printf("Start Server on :%d\n", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}

func NewServer(store kyverno.Store, port int, development bool) Server {
	return &server{
		store:       store,
		mux:         http.NewServeMux(),
		port:        port,
		development: development,
	}
}
