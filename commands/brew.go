package commands

//go:generate mockgen -source=brew.go -destination=mocks/brew.go

import (
	"os/exec"
	"strings"
)

type BrewCommander interface {
	RunBrewListFormula() ([]string, error)
	RunBrewListCask() ([]string, error)
	RunBrewInfoJSON(name string, isCask bool) ([]byte, error)
}

func NewBrewCommander() BrewCommander {
	return &brewCommander{}
}

type brewCommander struct{}

func (c *brewCommander) RunBrewListFormula() ([]string, error) {
	out, err := exec.Command("brew", "list", "--formula").Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(out)), "\n"), nil
}

func (c *brewCommander) RunBrewListCask() ([]string, error) {
	out, err := exec.Command("brew", "list", "--cask").Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(out)), "\n"), nil
}

func (c *brewCommander) RunBrewInfoJSON(name string, isCask bool) ([]byte, error) {
	args := []string{"info"}
	if isCask {
		args = append(args, "--cask")
	}
	args = append(args, "--json=v2", name)
	return exec.Command("brew", args...).Output()
}
