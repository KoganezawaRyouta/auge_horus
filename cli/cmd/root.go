package cmd

import (
	"fmt"

	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	importerCmd.Flags().StringVarP(&configName, "config", "c", "config", "default value is config")
	apiServerCmd.Flags().StringVarP(&configName, "config", "c", "config", "default value is config")
	appServerCmd.Flags().StringVarP(&configName, "config", "c", "config", "default value is config")
	importerCmd.Flags().StringVarP(&ENV, "env", "e", "development", "default value is development")
	apiServerCmd.Flags().StringVarP(&ENV, "env", "e", "development", "default value is development")
	appServerCmd.Flags().StringVarP(&ENV, "env", "e", "development", "default value is development")
	RootCmd.AddCommand(importerCmd)
	RootCmd.AddCommand(apiServerCmd)
	RootCmd.AddCommand(appServerCmd)
}

// LoadConfig db settings
func LoadConfig(configName string, ENV string) *settings.Config {
	var config *settings.Config
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName + "." + ENV)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("louding conf error: %s \n", err))
	}
	viper.Unmarshal(&config)
	return config
}
