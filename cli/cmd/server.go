package cmd

import (
	stdlog "log"
	"os"
	"strconv"

	"github.com/KoganezawaRyouta/augehorus/server"
	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

var PIDFile = "./tmp/api_server.pid"

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "auge horus web appication",
	Long:  `auge horus web appication`,
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

			setpidInfo(pidInfo, ppidInfo)
			removePIDFile()
			savePID(pid)
			config := LoadConfig(configName)
			errsCh <- server.ApiNew(config).Listen()
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

func removePIDFile() {
	err := os.Remove(PIDFile)
	if err != nil {
		stdlog.Printf("Unable to remove pid file : %v\n", err)
	}
}

func savePID(pid int) {

	file, err := os.Create(PIDFile)
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
