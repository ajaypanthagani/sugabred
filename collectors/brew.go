package collectors

//go:generate mockgen -source=brew.go -destination=mocks/brew.go

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ajaypanthagani/sugabred/commands"
	"github.com/ajaypanthagani/sugabred/types"
)

type BrewCollector interface {
	CollectPackages() ([]types.BrewPackage, error)
	CollectCasks() ([]types.BrewCask, error)
}

func NewBrewCollector(cmd commands.BrewCommander) BrewCollector {
	return &brewCollector{
		cmd: cmd,
	}
}

type brewCollector struct {
	cmd commands.BrewCommander
}

func (c *brewCollector) CollectPackages() ([]types.BrewPackage, error) {
	names, err := c.cmd.RunBrewListFormula()
	if err != nil {
		return nil, err
	}

	var pkgs []types.BrewPackage

	for _, name := range names {
		if strings.TrimSpace(name) == "" {
			continue
		}

		fmt.Printf("Processing package: %s\n", name)

		out, err := c.cmd.RunBrewInfoJSON(name, false)
		if err != nil {
			continue
		}

		var info types.BrewPackageInfo
		if err := json.Unmarshal(out, &info); err != nil || len(info.Formulae) == 0 {
			continue
		}

		pkgs = append(pkgs, types.BrewPackage{
			Name:    info.Formulae[0].Name,
			Version: info.Formulae[0].Versions.Stable,
		})
	}
	return pkgs, nil
}

func (c *brewCollector) CollectCasks() ([]types.BrewCask, error) {
	names, err := c.cmd.RunBrewListCask()

	if err != nil {
		return nil, err
	}

	var casks []types.BrewCask

	for _, token := range names {
		if strings.TrimSpace(token) == "" {
			continue
		}

		fmt.Printf("Processing cask: %s\n", token)

		out, err := c.cmd.RunBrewInfoJSON(token, true)
		if err != nil {
			continue
		}

		var info types.BrewCaskInfo
		if err := json.Unmarshal(out, &info); err != nil || len(info.Casks) == 0 {
			continue
		}

		casks = append(casks, types.BrewCask{
			Name:    info.Casks[0].Token,
			Version: info.Casks[0].Version,
		})
	}

	return casks, nil
}
