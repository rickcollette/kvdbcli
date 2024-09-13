package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/rickcollette/kayveedb"
)

var (
	snapshot       = "path/to/snapshot"
	logPath        = "path/to/log"
	hmacKey        = []byte("32-byte-hmac-key")
	encryptionKey  = []byte("32-byte-long-encryption-key")
	nonce          = []byte("24-byte-nonce")
	keySize        = 32 // XChaCha20 key size
	nonceSize      = 24 // XChaCha20 nonce size
)

// validateEncryptionParams validates the encryption key and nonce.
func validateEncryptionParams() error {
	if len(encryptionKey) != keySize {
		return fmt.Errorf("encryptionKey must be %d bytes long", keySize)
	}
	if len(nonce) != nonceSize {
		return fmt.Errorf("nonce must be %d bytes long", nonceSize)
	}
	return nil
}

// loadBTree loads the BTree from the snapshot and logs.
func loadBTree() (*kayveedb.BTree, error) {
	if err := validateEncryptionParams(); err != nil {
		return nil, err
	}
	btree, err := kayveedb.NewBTree(3, snapshot, logPath, hmacKey, encryptionKey, nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to load B-tree: %v", err)
	}
	return btree, nil
}

// snapshotBTree takes a snapshot of the current BTree state.
func snapshotBTree(btree *kayveedb.BTree) error {
	if err := btree.Snapshot(); err != nil {
		return fmt.Errorf("failed to take snapshot: %v", err)
	}
	log.Println("Snapshot taken successfully")
	return nil
}

// deleteKey deletes a key from the BTree.
func deleteKey(btree *kayveedb.BTree, key string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	if err := btree.Delete(key); err != nil {
		return fmt.Errorf("failed to delete key %s: %v", key, err)
	}
	log.Printf("Deleted key: %s\n", key)
	return nil
}

// insertKey inserts a new key-value pair into the BTree.
func insertKey(btree *kayveedb.BTree, key, value string) error {
	if key == "" || value == "" {
		return errors.New("key and value cannot be empty")
	}
	encValue := []byte(value)
	if err := btree.Insert(key, encValue, encryptionKey, nonce); err != nil {
		return fmt.Errorf("failed to insert key %s: %v", key, err)
	}
	log.Printf("Inserted key: %s\n", key)
	return nil
}

// updateKey updates an existing key-value pair in the BTree.
func updateKey(btree *kayveedb.BTree, key, value string) error {
	if key == "" || value == "" {
		return errors.New("key and value cannot be empty")
	}
	encValue := []byte(value)
	if err := btree.Update(key, encValue, encryptionKey, nonce); err != nil {
		return fmt.Errorf("failed to update key %s: %v", key, err)
	}
	log.Printf("Updated key: %s\n", key)
	return nil
}

// readKey reads the value of a key from the BTree.
func readKey(btree *kayveedb.BTree, key string) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}
	value, err := btree.Read(key, encryptionKey, nonce)
	if err != nil {
		return "", fmt.Errorf("failed to read key %s: %v", key, err)
	}
	return string(value), nil
}
