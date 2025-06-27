/*
Copyright © 2025 Ajay Panthagani ajaypanthagani321@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sugabred",
	Short: "Sugabred is a macOS tool that snapshots your entire developer environment — from tools to settings — and makes it effortlessly reproducible on another Mac.",
	Long: `Sugabred is a foolproof, cross-architecture developer environment replication tool for macOS. 
	It allows you to capture the entire state of your development machine — including system settings, tools, packages, configurations, environment variables, plugins, and more — and reproduce it on another Mac with absolute fidelity.
	
	Unlike dotfile managers or configuration sync tools, Sugabred goes deeper:
	- Captures both CLI and GUI software with exact versions
	- Records Homebrew packages, shell setups (zsh, bash, fish), system defaults, IDE settings, browser extensions, LaunchAgents, fonts, VPN profiles, network configs, and developer caches (like .ivy2, .m2, .nvm, etc.)
	- Intelligently maps installations across architectures (e.g., ARM ↔ x86) when restoring on a different Apple chip
	- Excludes personal files (like documents, photos, private code) while preserving a clean, reproducible development setup
	
	It is designed to make onboarding new developers instant, migrating to a new Mac trivial, and CI or fleet setups deterministic — down to the last hidden config file.
	
	Sugabred gives developers superpowers by ensuring: 
	"If it runs on my machine, it will run on yours — identically."
	
	Core commands:
	  - snapshot: Save your current machine’s dev setup into a portable YAML file
	  - up: Recreate the entire dev environment from a Sugabred snapshot
	  - doctor: Diagnose mismatches between a snapshot and the current machine
	
	Sugabred is open source, built in Go, and made for engineers who care deeply about reproducibility, environment parity, and development velocity.
	
	Use Sugabred to eliminate “it works on my machine” forever.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
