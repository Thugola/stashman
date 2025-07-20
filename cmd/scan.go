package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Test")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
