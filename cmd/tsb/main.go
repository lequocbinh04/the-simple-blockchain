package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var tsbCmd = &cobra.Command{
		Use:   "tsb",
		Short: "The Simple Blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	tsbCmd.AddCommand(versionCmd)
	tsbCmd.AddCommand(balancesCmd())
	tsbCmd.AddCommand(txCmd())

	err := tsbCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
