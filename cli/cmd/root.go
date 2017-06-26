package cmd

import (
	"github.com/spf13/cobra"
)

var configName string
var ENV string
var RootCmd = &cobra.Command{
	Use:   "augehorus",
	Short: "short description",
	Long:  `long description`,
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(importerCmd)
	RootCmd.AddCommand(apiServerCmd)
	RootCmd.AddCommand(appServerCmd)
}
