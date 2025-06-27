/*
Copyright © 2025 Ajay Panthagani ajaypanthagani321@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/ajaypanthagani/sugabred/collectors"
	"github.com/ajaypanthagani/sugabred/commands"
	"github.com/ajaypanthagani/sugabred/output"
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Creates a complete YAML snapshot of the current macOS developer environment, including tools, configs, system state, and more — excluding personal files.",
	Long: `The 'snapshot' command inspects your current macOS machine and generates a detailed YAML file capturing your entire developer setup.

	It includes:
	- Installed applications and versions (CLI + GUI)
	- Homebrew packages and taps
	- System-level configurations (e.g., defaults)
	- Shell environment and dotfiles
	- IDEs, plugins, and editor settings
	- Environment variables
	- VPN profiles and network settings
	- Installed fonts, launch agents, and dev-specific caches (like .ivy2)
	- Git repositories in workspace folders

	It explicitly excludes personal user files, sensitive documents, and non-dev data. You can customize what’s included/excluded using flags or config overrides.

	Use this command to create a portable, versionable snapshot of your dev environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Collecting snapshot...")
		snap, err := devEnvCollector.CollectAll()
		if err != nil {
			fmt.Println("❌ Failed to generate snapshot: %w", err)
			return
		}

		err = output.WriteSnapshotToFile(snap, "sugabred.snapshot.yaml")
		if err != nil {
			fmt.Println("❌ Failed to write snapshot: %w", err)
			return
		}

		fmt.Println("✅ Snapshot written to sugabred.snapshot.yaml")
	},
}

var (
	brewCommander   commands.BrewCommander
	envCommander    commands.EnvCommander
	brewCollector   collectors.BrewCollector
	envCollector    collectors.EnvCollector
	devEnvCollector collectors.DevEnvCollector
)

func setup() {
	brewCommander = commands.NewBrewCommander()
	envCommander = commands.NewEnvCommander()

	brewCollector = collectors.NewBrewCollector(brewCommander)
	envCollector = collectors.NewEnvCollector(envCommander)
	devEnvCollector = collectors.NewDevEnvCollector(brewCollector, envCollector)
}

func init() {
	setup()
	rootCmd.AddCommand(snapshotCmd)
}
