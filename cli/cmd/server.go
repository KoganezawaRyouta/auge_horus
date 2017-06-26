package cmd

import (
	stdlog "log"
	"os"

	webApi "github.com/KoganezawaRyouta/augehorus/api"
	webApp "github.com/KoganezawaRyouta/augehorus/app"
	"github.com/KoganezawaRyouta/augehorus/config"
	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

var apiServerCmd = &cobra.Command{
	Use:   "api_server",
	Short: "auge horus api server",
	Long:  `auge horus api server`,
	Run: func(cmd *cobra.Command, args []string) {
		errsCh := make(chan error)
		go func() {
			pid := os.Getpid()
			pidInfo, _ := ps.FindProcess(pid)
			stdlog.Printf("start server!!")
			stdlog.Printf(" PID          : %d\n", pidInfo.Pid())
			stdlog.Printf(" PPID         : %d\n", pidInfo.PPid())
			stdlog.Printf(" Process name : %s\n", pidInfo.Executable())
			ppidInfo, _ := ps.FindProcess(pidInfo.PPid())
			stdlog.Printf("\n=================")
			stdlog.Printf(" Parent process name : %s\n", ppidInfo.Executable())

			config := config.ConfigNew()
			stdlog.Printf(" Port          : %d\n", config.ApiServer.Port)
			errsCh <- webApi.ApiNew(config).Listen()
		}()
		stdlog.Fatalln("terminated", <-errsCh)
	},
}

var appServerCmd = &cobra.Command{
	Use:   "app_server",
	Short: "auge horus app server",
	Long:  `auge horus app server`,
	Run: func(cmd *cobra.Command, args []string) {
		errsCh := make(chan error)
		go func() {
			pid := os.Getpid()
			pidInfo, _ := ps.FindProcess(pid)
			stdlog.Printf("start server!!")
			stdlog.Printf(" PID          : %d\n", pidInfo.Pid())
			stdlog.Printf(" PPID         : %d\n", pidInfo.PPid())
			stdlog.Printf(" Process name : %s\n", pidInfo.Executable())
			ppidInfo, _ := ps.FindProcess(pidInfo.PPid())
			stdlog.Printf("\n=================")
			stdlog.Printf(" Parent process name : %s\n", ppidInfo.Executable())

			config := config.ConfigNew()
			stdlog.Printf(" Port          : %d\n", config.ApiServer.Port)
			errsCh <- webApp.AppNew(config).Listen()
		}()
		stdlog.Fatalln("terminated", <-errsCh)
	},
}
