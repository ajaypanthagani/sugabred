package collectors

//go:generate mockgen -source=collector.go -destination=mocks/collector.go

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/ajaypanthagani/sugabred/types"
)

type DevEnvCollector interface {
	CollectAll() (*types.Snapshot, error)
}

func NewDevEnvCollector(brewCollector BrewCollector, envCollector EnvCollector) DevEnvCollector {
	return &devEnvCollector{
		brewCollector: brewCollector,
		envCollector:  envCollector,
	}
}

type devEnvCollector struct {
	brewCollector BrewCollector
	envCollector  EnvCollector
}

func (c *devEnvCollector) CollectAll() (*types.Snapshot, error) {
	fmt.Println("Collecting homebrew packages...")
	brewPkgs, err := c.brewCollector.CollectPackages()
	if err != nil {
		return nil, err
	}

	fmt.Println("Collecting homebrew casks...")
	brewCasks, err := c.brewCollector.CollectCasks()
	if err != nil {
		return nil, err
	}

	fmt.Println("Collecting env variables...")
	envVars := c.envCollector.CollectEnvVars()

	timestamp := time.Now().Format(time.RFC3339)
	architecture := runtime.GOARCH
	macosVersion, err := getMacOSVersion()
	if err != nil {
		macosVersion = "unknown"
	}

	return &types.Snapshot{
		Timestamp: timestamp,
		Arch:      architecture,
		OSVersion: macosVersion,
		Homebrew:  brewPkgs,
		Casks:     brewCasks,
		EnvVars:   envVars,
	}, nil
}

func getMacOSVersion() (string, error) {
	out, err := exec.Command("sw_vers", "-productVersion").Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
