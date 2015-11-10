package runtime

import (
	"runtime"
)

// AllProcs configure runtime to use all available processors/cpus.
func AllProcs() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
