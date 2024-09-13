package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	key   string
	value string
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert a key-value pair into KayveeDB",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := LoadBtree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}

		// Use helper function to insert the key
		if err := InsertKey(btree, key, value); err != nil {
			log.Fatalf("Error inserting key: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringVarP(&key, "key", "k", "", "Key to insert")
	insertCmd.Flags().StringVarP(&value, "value", "v", "", "Value to insert")
	insertCmd.MarkFlagRequired("key")
	insertCmd.MarkFlagRequired("value")
}
