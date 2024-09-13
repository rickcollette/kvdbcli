package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a key-value pair from KayveeDB",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the B-tree using the helper function
		btree, err := LoadBtree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}

		// Use the deleteKey helper function for deletion
		err = DeleteKey(btree, key)
		if err != nil {
			log.Fatalf("Error deleting key %s: %v", key, err)
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	// Set up the flag for the key to delete
	deleteCmd.Flags().StringVarP(&key, "key", "k", "", "Key to delete")
	// Mark the key flag as required
	deleteCmd.MarkFlagRequired("key")
}
