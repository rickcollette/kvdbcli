package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"kayveedb"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Create a snapshot of the current KayveeDB state",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := loadBTree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}

		err = btree.Snapshot()
		if err != nil {
			log.Fatalf("Failed to create snapshot: %v", err)
		}
		fmt.Println("Snapshot created successfully")
	},
}

func init() {
	rootCmd.AddCommand(snapshotCmd)
}
