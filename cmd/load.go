package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load the KayveeDB from logs",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := loadBTree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}
		// Using the loaded btree or processing further as needed
		if btree != nil {
			log.Println("KayveeDB loaded from logs successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
