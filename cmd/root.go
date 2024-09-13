package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kayveedb",
	Short: "KayveeDB CLI is a tool to interact with the KayveeDB B-tree encrypted key-value store",
	Long: `KayveeDB CLI allows users to perform key-value operations like inserting, reading, updating, and deleting,
along with managing snapshots and logs of the B-tree database.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
