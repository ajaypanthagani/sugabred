/*
Copyright © 2025 Ajay Panthagani ajaypanthagani321@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Reproduces the environment from a Sugabred YAML snapshot on a fresh macOS machine.",
	Long: `The 'up' command reads a Sugabred snapshot YAML file and reproduces the exact developer environment on a macOS machine.

	This includes installing system packages, developer tools, CLI utilities, configuration files, environment variables, network profiles, IDEs, shell settings, browser extensions, and more — all based on the snapshot.
	
	It intelligently handles architecture differences (e.g., ARM vs x86) and ensures compatible versions are installed.

	Use this to set up a new machine or onboard a teammate with a guaranteed identical dev environment.
	
	By default, Sugabred looks for a 'sugabred.snapshot.yaml' in the current directory, but you can pass a custom path using the --file flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("up called")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
