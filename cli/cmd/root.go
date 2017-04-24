package cmd

import (
	"fmt"

	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configName string
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

// LoadConfig db settings
func LoadConfig(configName string) *settings.Config {
	var config *settings.Config
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("louding conf error: %s \n", err))
	}
	viper.Unmarshal(&config)
	return config
}
