package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Create a snapshot of the current KayveeDB state",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := loadBTree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}

		// Use helper function to take a snapshot
		if err := snapshotBTree(btree); err != nil {
			log.Fatalf("Failed to take snapshot: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(snapshotCmd)
}
