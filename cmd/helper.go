package cmd

import (
	"fmt"
	"log"
	"os"
	"kayveedb"
)

func loadBTree() (*kayveedb.BTree, error) {
	snapshot := "path/to/snapshot"
	logPath := "path/to/log"
	hmacKey := []byte("32-byte-hmac-key")
	encryptionKey := []byte("32-byte-long-encryption-key")
	nonce := []byte("24-byte-nonce")

	btree, err := kayveedb.NewBTree(3, snapshot, logPath, hmacKey, encryptionKey, nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to load B-tree: %v", err)
	}

	return btree, nil
}
