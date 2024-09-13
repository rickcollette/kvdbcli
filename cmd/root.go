package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "kvdbcli",
	Short: "kvdbcli is a tool to interact with the KayveeDB B-tree encrypted key-value store",
	Long: `kvdbcli allows users to perform key-value operations like inserting, reading, updating, and deleting,
along with managing snapshots and logs of the B-tree database.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		LoadConfig()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Flags for snapshot and log path
	RootCmd.PersistentFlags().StringVar(&Snapshot, "snapshot", "/path/to/snapshot", "Path to the snapshot file")
	RootCmd.PersistentFlags().StringVar(&LogPath, "logpath", "/path/to/log", "Path to the log file")

	// Bind flags to viper
	viper.BindPFlag("snapshot", RootCmd.PersistentFlags().Lookup("snapshot"))
	viper.BindPFlag("logpath", RootCmd.PersistentFlags().Lookup("logpath"))
}
