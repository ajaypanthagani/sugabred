package commands

//go:generate mockgen -source=env.go -destination=mocks/env.go

import "os"

type EnvCommander interface {
	Environ() []string
}

func NewEnvCommander() EnvCommander {
	return &envCommander{}
}

type envCommander struct{}

func (c *envCommander) Environ() []string {
	return os.Environ()
}
