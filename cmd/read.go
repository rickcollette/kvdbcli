package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a value by key from KayveeDB",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := LoadBtree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}

		// Use helper function to read the key
		val, err := ReadKey(btree, key)
		if err != nil {
			log.Fatalf("Error reading key: %v", err)
		}
		log.Printf("Value for key '%s': %s\n", key, val)
	},
}

func init() {
	RootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&key, "key", "k", "", "Key to read")
	readCmd.MarkFlagRequired("key")
}
