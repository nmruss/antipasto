package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number of Antipasto",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Antipasto digital asset generator v0.1")
	},
}
