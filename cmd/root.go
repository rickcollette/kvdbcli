package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "kayveedb",
	Short: "KayveeDB CLI is a tool to interact with the KayveeDB B-tree encrypted key-value store",
	Long: `KayveeDB CLI allows users to perform key-value operations like inserting, reading, updating, and deleting,
along with managing snapshots and logs of the B-tree database.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		LoadConfig()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Flags for snapshot and log path
	rootCmd.PersistentFlags().StringVar(&Snapshot, "snapshot", "/path/to/snapshot", "Path to the snapshot file")
	rootCmd.PersistentFlags().StringVar(&LogPath, "logpath", "/path/to/log", "Path to the log file")

	// Bind flags to viper
	viper.BindPFlag("snapshot", rootCmd.PersistentFlags().Lookup("snapshot"))
	viper.BindPFlag("logpath", rootCmd.PersistentFlags().Lookup("logpath"))
}
