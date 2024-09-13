package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var SetCmd = &cobra.Command{
	Use:   "set KEY=VALUE",
	Short: "Set environment variables or config values",
	Long: `The 'set' command allows you to set environment variables or configuration values. 
Usage:
  set HMAC=your-hmac-key
  set ENCRYPTION=your-encryption-key`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the user provided the correct format
		if len(args) != 1 {
			printSetUsage()
			return
		}

		// Handle the key-value setting
		handleSetCommand(args[0])
	},
}

// handleSetCommand processes the input for setting configuration
func handleSetCommand(setCmd string) {
	parts := strings.SplitN(setCmd, "=", 2)
	if len(parts) != 2 {
		printSetUsage()
		return
	}

	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	// Update the value in the config or environment
	if viper.IsSet(key) || validKeys[key] {
		viper.Set(key, value)
		fmt.Printf("%s set to %s\n", key, value)
	} else {
		fmt.Printf("Unknown configuration key: %s\n", key)
		printSetUsage()
	}
}

// printSetUsage shows the correct usage for the 'set' command
func printSetUsage() {
	fmt.Println("Usage: set KEY=VALUE")
	fmt.Println("Valid keys: HMAC, ENCRYPTION, NONCE, snapshot, logpath")
}

var validKeys = map[string]bool{
	"HMAC":       true,
	"ENCRYPTION": true,
	"NONCE":      true,
	"snapshot":   true,
	"logpath":    true,
}

func init() {
	RootCmd.AddCommand(SetCmd)
}
