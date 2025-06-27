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
		fmt.Printf("Error retrieving brew packages: %s", err.Error())
		return nil, err
	}

	var pkgs []types.BrewPackage

	for _, name := range names {
		if strings.TrimSpace(name) == "" {
			continue
		}

		fmt.Printf("Processing brew package: %s\n", name)

		out, err := c.cmd.RunBrewInfoJSON(name, false)
		if err != nil {
			fmt.Printf("Couldn't process brew package: %s, received err: %s", name, err.Error())
			continue
		}

		var info types.BrewPackageInfo
		if err := json.Unmarshal(out, &info); err != nil {
			fmt.Printf("Couldn't unmarshal brew package info: %s for package: %s, received err: %s", out, name, err.Error())
			continue
		}

		if len(info.Formulae) == 0 {
			fmt.Printf("Couldn't process brew package: %s, received empty formula list", name)
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
		fmt.Printf("Error retrieving brew casks: %s", err.Error())
		return nil, err
	}

	var casks []types.BrewCask

	for _, token := range names {
		if strings.TrimSpace(token) == "" {
			continue
		}

		fmt.Printf("Processing brew cask: %s\n", token)

		out, err := c.cmd.RunBrewInfoJSON(token, true)
		if err != nil {
			fmt.Printf("Couldn't process brew cask: %s, received err: %s", token, err.Error())
			continue
		}

		var info types.BrewCaskInfo
		if err := json.Unmarshal(out, &info); err != nil {
			fmt.Printf("Couldn't unmarshal brew cask package info: %s for cask: %s, received err: %s", out, token, err.Error())
			continue
		}

		if len(info.Casks) == 0 {
			fmt.Printf("Couldn't process brew cask: %s, received empty formula list", token)
			continue
		}

		casks = append(casks, types.BrewCask{
			Name:    info.Casks[0].Token,
			Version: info.Casks[0].Version,
		})
	}

	return casks, nil
}
