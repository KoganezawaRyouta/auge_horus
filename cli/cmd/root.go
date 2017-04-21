package cmd

import (
	"github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
)

var configName string
var logger log.Logger
var RootCmd = &cobra.Command{
	Use:   "batch coint app",
	Short: "short description",
	Long:  `long description`,
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}

func init() {
	cobra.OnInitialize()
	importerCmd.Flags().StringVarP(&configName, "config", "c", "development_config", "default value is development_config")
	serverCmd.Flags().StringVarP(&configName, "config", "c", "development_config", "default value is development_config")
	RootCmd.AddCommand(importerCmd)
	RootCmd.AddCommand(serverCmd)
}
