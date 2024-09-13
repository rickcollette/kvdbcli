package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"kayveedb"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a value by key in KayveeDB",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := loadBTree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}

		encryptionKey := []byte("32-byte-long-encryption-key")
		nonce := []byte("24-byte-nonce")
		err = btree.Update(key, []byte(value), encryptionKey, nonce)
		if err != nil {
			log.Fatalf("Failed to update key-value: %v", err)
		}
		fmt.Printf("Updated key: %s\n", key)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&key, "key", "k", "", "Key to update")
	updateCmd.Flags().StringVarP(&value, "value", "v", "", "New value to update")
	updateCmd.MarkFlagRequired("key")
	updateCmd.MarkFlagRequired("value")
}
