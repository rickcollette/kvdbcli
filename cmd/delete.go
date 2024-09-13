package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"kayveedb"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a key-value pair from KayveeDB",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := loadBTree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}

		err = btree.Delete(key)
		if err != nil {
			log.Fatalf("Failed to delete key: %v", err)
		}
		fmt.Printf("Deleted key: %s\n", key)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&key, "key", "k", "", "Key to delete")
	deleteCmd.MarkFlagRequired("key")
}
