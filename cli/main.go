package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KoganezawaRyouta/augehorus/cli/cmd"
	"github.com/KoganezawaRyouta/augehorus/config"
	"github.com/spf13/cobra"
)

func main() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  `Print the version number`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("batch v%s-%s\n", config.Version, config.GoVersion)
			log.Printf("BuildDhash%s\n", config.BuildDhash)
		},
	}
	cmd.RootCmd.AddCommand(versionCmd)

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
