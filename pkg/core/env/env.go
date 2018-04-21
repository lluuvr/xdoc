package env

import (
	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/core/runtime"
)

var (
	e = &env{}
)

type env struct {
	broker  *broker.Broker
	runtime *runtime.Runtime
}

// SetRuntime set runtime to env
func SetRuntime(rt *runtime.Runtime) {
	e.runtime = rt
}

// GetRuntime get runtime from env
func GetRuntime() *runtime.Runtime {
	return e.runtime
}

// SetBroker set broker to env
func SetBroker(brk *broker.Broker) {
	e.broker = brk
}

// GetBroker returns broker from env
func GetBroker() *broker.Broker {
	return e.broker
}
