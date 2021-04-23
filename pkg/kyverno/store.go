package kyverno

import "sync"

type Store interface {
	GetMetrics() Metrics
	SetMetrics(Metrics)
}

type store struct {
	mx *sync.RWMutex
	m  Metrics
}

func (s *store) GetMetrics() Metrics {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.m
}

func (s *store) SetMetrics(m Metrics) {
	s.mx.Lock()
	s.m = m
	s.mx.Unlock()
}

func NewStore() Store {
	return &store{
		mx: new(sync.RWMutex),
	}
}
