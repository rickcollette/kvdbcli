package cmd

import (
	"errors"
	"fmt"
	"log"
	"github.com/spf13/viper"
	"github.com/rickcollette/kayveedb"
)

var (
	Snapshot      = "path/to/snapshot"
	LogPath       = "path/to/log"
	HmacKey       = []byte("32-byte-hmac-key")
	EncryptionKey = []byte("32-byte-long-encryption-key")
	Nonce         = []byte("24-byte-nonce")
	KeySize       = 32  // XChaCha20 key size
	NonceSize     = 24  // XChaCha20 nonce size
	CacheSize     = 100 // Set a default cache size for B-tree nodes
)

// LoadConfig loads the config file and environment variables using Viper
func LoadConfig() {
	viper.SetConfigName("kayveedb")           // Config file name without extension
	viper.AddConfigPath("/etc/kayveedb/")     // Look for config in /etc
	viper.AddConfigPath("$HOME/.kayveedb/")   // Look for config in the home directory
	viper.AddConfigPath(".")                  // Look for config in the current directory

	// Bind environment variables
	viper.BindEnv("HMAC_KEY")
	viper.BindEnv("ENCRYPTION_KEY")
	viper.BindEnv("NONCE")

	// Try to read the config file
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("No config file found. Using default values or environment variables.")
	}
}
// InitializeKeys initializes the keys from environment variables or errors out.
func InitializeKeys() error {
	HmacKey = []byte(viper.GetString("HMAC_KEY"))
	EncryptionKey = []byte(viper.GetString("ENCRYPTION_KEY"))
	Nonce = []byte(viper.GetString("NONCE"))

	if len(HmacKey) != KeySize || len(EncryptionKey) != KeySize || len(Nonce) != NonceSize {
		return fmt.Errorf("invalid key lengths: hmacKey and encryptionKey must be %d bytes, nonce must be %d bytes", KeySize, NonceSize)
	}
	return nil
}

// ValidateEncryptionParams validates the encryption key and nonce.
func ValidateEncryptionParams() error {
	if len(EncryptionKey) != KeySize {
		return fmt.Errorf("encryptionKey must be %d bytes long", KeySize)
	}
	if len(Nonce) != NonceSize {
		return fmt.Errorf("nonce must be %d bytes long", NonceSize)
	}
	return nil
}

// LoadBtree loads the B-tree from the snapshot and logs.
func LoadBtree() (*kayveedb.BTree, error) {
	if err := ValidateEncryptionParams(); err != nil {
		return nil, err
	}
	btree, err := kayveedb.NewBTree(3, Snapshot, LogPath, HmacKey, EncryptionKey, Nonce, CacheSize)
	if err != nil {
		return nil, fmt.Errorf("failed to load B-tree: %v", err)
	}
	return btree, nil
}

// SnapshotBtree takes a snapshot of the current B-tree state.
func SnapshotBtree(btree *kayveedb.BTree) error {
	if err := btree.Snapshot(); err != nil {
		return fmt.Errorf("failed to take snapshot: %v", err)
	}
	log.Println("Snapshot taken successfully")
	return nil
}

// DeleteKey deletes a key from the B-tree.
func DeleteKey(btree *kayveedb.BTree, key string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	if err := btree.Delete(nil, key); err != nil {
		return fmt.Errorf("failed to delete key %s: %v", key, err)
	}
	log.Printf("Deleted key: %s\n", key)
	return nil
}

// InsertKey inserts a new key-value pair into the B-tree.
func InsertKey(btree *kayveedb.BTree, key, value string) error {
	if key == "" || value == "" {
		return errors.New("key and value cannot be empty")
	}
	encValue := []byte(value)
	if err := btree.Insert(key, encValue, EncryptionKey, Nonce); err != nil {
		return fmt.Errorf("failed to insert key %s: %v", key, err)
	}
	log.Printf("Inserted key: %s\n", key)
	return nil
}

// UpdateKey updates an existing key-value pair in the B-tree.
func UpdateKey(btree *kayveedb.BTree, key, value string) error {
	if key == "" || value == "" {
		return errors.New("key and value cannot be empty")
	}
	encValue := []byte(value)
	if err := btree.Update(key, encValue, EncryptionKey, Nonce); err != nil {
		return fmt.Errorf("failed to update key %s: %v", key, err)
	}
	log.Printf("Updated key: %s\n", key)
	return nil
}

// ReadKey reads the value of a key from the B-tree.
func ReadKey(btree *kayveedb.BTree, key string) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}
	value, err := btree.Read(key, EncryptionKey, Nonce)
	if err != nil {
		return "", fmt.Errorf("failed to read key %s: %v", key, err)
	}
	return string(value), nil
}
