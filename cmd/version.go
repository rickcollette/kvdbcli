package cmd

import (
	"fmt"

	"github.com/rickcollette/kayveedb"
	"github.com/spf13/cobra"
)

const Version string = "v1.0.5"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of kvdbcli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kvdbcli version: %s\nkayveedb version: %s\n", Version, kayveedb.ShowVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
