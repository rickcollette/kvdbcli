package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a value by key in KayveeDB",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := LoadBtree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}

		// Use helper function to update the key
		if err := UpdateKey(btree, key, value); err != nil {
			log.Fatalf("Error updating key: %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&key, "key", "k", "", "Key to update")
	updateCmd.Flags().StringVarP(&value, "value", "v", "", "New value to update")
	updateCmd.MarkFlagRequired("key")
	updateCmd.MarkFlagRequired("value")
}
