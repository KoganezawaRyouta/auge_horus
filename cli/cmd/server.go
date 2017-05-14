package cmd

import (
	stdlog "log"
	"os"
	"strconv"

	webApp "github.com/KoganezawaRyouta/augehorus/app"
	webApi "github.com/KoganezawaRyouta/augehorus/api"
	"github.com/KoganezawaRyouta/augehorus/settings"
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

			config := LoadConfig(configName, ENV)
			setpidInfo(pidInfo, ppidInfo)
			removePIDFile(config.ApiServer.PidFile)
			savePID(config.ApiServer.PidFile, pid)
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

			config := LoadConfig(configName, ENV)
			setpidInfo(pidInfo, ppidInfo)
			removePIDFile(config.AppServer.PidFile)
			savePID(config.AppServer.PidFile, pid)
			errsCh <- webApp.AppNew(config).Listen()
		}()
		stdlog.Fatalln("terminated", <-errsCh)
	},
}

func setpidInfo(pidInfo ps.Process, ppidInfo ps.Process) {
	settings.PID = pidInfo.Pid()
	settings.PPID = pidInfo.PPid()
	settings.ProcessName = pidInfo.Executable()
	settings.ParentProcessName = ppidInfo.Executable()
}

func removePIDFile(pidFile string) {
	err := os.Remove(pidFile)
	if err != nil {
		stdlog.Printf("Unable to remove pid file : %v\n", err)
	}
}

func savePID(pidFile string, pid int) {

	file, err := os.Create(pidFile)
	if err != nil {
		stdlog.Printf("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))
	if err != nil {
		stdlog.Printf("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	file.Sync() // flush to disk
}
