package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/rickcollette/kayveedb"
)

var (
	key   string
	value string
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert a key-value pair into KayveeDB",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := loadBTree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}

		encryptionKey := []byte("32-byte-long-encryption-key")
		nonce := []byte("24-byte-nonce")
		err = btree.Insert(key, []byte(value), encryptionKey, nonce)
		if err != nil {
			log.Fatalf("Failed to insert key-value: %v", err)
		}
		fmt.Printf("Inserted key: %s\n", key)
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringVarP(&key, "key", "k", "", "Key to insert")
	insertCmd.Flags().StringVarP(&value, "value", "v", "", "Value to insert")
	insertCmd.MarkFlagRequired("key")
	insertCmd.MarkFlagRequired("value")
}
