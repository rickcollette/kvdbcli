package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"kayveedb"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load the KayveeDB from logs",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := loadBTree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}
		fmt.Println("Loaded KayveeDB from logs successfully")
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
