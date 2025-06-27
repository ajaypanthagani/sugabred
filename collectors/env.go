package collectors

//go:generate mockgen -source=env.go -destination=mocks/env.go

import (
	"strings"

	"github.com/ajaypanthagani/sugabred/commands"
)

type EnvCollector interface {
	CollectEnvVars() map[string]string
}

func NewEnvCollector(envCommander commands.EnvCommander) EnvCollector {
	return &envCollector{
		envCommander: envCommander,
	}
}

type envCollector struct {
	envCommander commands.EnvCommander
}

func (c *envCollector) CollectEnvVars() map[string]string {
	envMap := make(map[string]string)
	for _, kv := range c.envCommander.Environ() {
		parts := strings.SplitN(kv, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}
	return envMap
}
