/*
Copyright © 2025 Ajay Panthagani ajaypanthagani321@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Diagnoses environment mismatches between the current machine and a Sugabred snapshot, and suggests exact steps to fix them.",
	Long: `The 'doctor' command compares the current machine's environment with a Sugabred snapshot and diagnoses discrepancies.

	It highlights:
	- Missing or mismatched software versions
	- Uninstalled tools or system configurations
	- Diverging environment variables
	- Conflicting settings or missing dependencies
	- Architecture mismatches (e.g., x86_64 vs arm64 binaries)

	For each issue, it suggests actionable steps to reconcile the difference and align the machine with the snapshot.

	This is useful for debugging inconsistencies between dev machines, CI environments, or restoring parity after drift.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("doctor called")
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
