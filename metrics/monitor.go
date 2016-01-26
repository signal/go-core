// Meters and stats endpoint for Go projects
package metrics

import (
	"gopkg.in/signal/go-metrics.v1"
	"sync"
)

func Counter(target string) metrics.Counter {
	mux.Lock()
	defer mux.Unlock()
	if _, exists := counters[target]; !exists {
		meter := metrics.NewCounter()
		metrics.Register(target, meter)
		counters[target] = meter
	}
	return counters[target]
}

func Registry() metrics.Registry {
	return metrics.DefaultRegistry
}

/* private below here */

var (
	counters map[string]metrics.Counter
	mux = &sync.Mutex{}
)

func init() {
	counters = make(map[string]metrics.Counter)
}

