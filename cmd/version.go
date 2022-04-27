/*
Copyright © 2022 Dan Anstis

*/
package cmd

import (
	"fmt"

	"github.com/danstis/taggr/internal/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the application version",
	Long: `Used to check the particular version of taggr that is installed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("taggr v%s\n", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
