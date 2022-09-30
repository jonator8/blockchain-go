package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "bc",
		Short: "The Blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(balancesCmd())
	cmd.AddCommand(txCmd())

	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
