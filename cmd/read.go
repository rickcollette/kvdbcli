package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"kayveedb"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a value by key from KayveeDB",
	Run: func(cmd *cobra.Command, args []string) {
		btree, err := loadBTree()
		if err != nil {
			log.Fatalf("Failed to load B-tree: %v", err)
		}

		encryptionKey := []byte("32-byte-long-encryption-key")
		nonce := []byte("24-byte-nonce")
		value, err := btree.Read(key, encryptionKey, nonce)
		if err != nil {
			log.Fatalf("Failed to read key: %v", err)
		}
		fmt.Printf("Value for key %s: %s\n", key, string(value))
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&key, "key", "k", "", "Key to read")
	readCmd.MarkFlagRequired("key")
}
