/*
Copyright © 2025 Ajay Panthagani ajaypanthagani321@gmail.com
*/
package cmd

import (
	"fmt"

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
		fmt.Println("snapshot called")
	},
}

func init() {
	rootCmd.AddCommand(snapshotCmd)
}
