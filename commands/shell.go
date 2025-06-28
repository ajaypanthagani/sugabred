package commands

//go:generate mockgen -source=shell.go -destination=mocks/shell.go

import (
	"os/exec"
)

type ShellCommander interface {
	RunCommand(name string, args ...string) (string, error)
}

func NewShellCommander() ShellCommander {
	return &shellCommander{}
}

type shellCommander struct {
}

func (*shellCommander) RunCommand(name string, args ...string) (string, error) {
	out, err := exec.Command(name, args...).CombinedOutput()
	return string(out), err
}
