package cmd

import (
	stdlog "log"
	"os"
	"strconv"

	"github.com/KoganezawaRyouta/augehorus/server"
	"github.com/spf13/cobra"
)

var PIDFile = "./tmp/api_server.pid"

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "bitcoin rate infomation web appication",
	Long:  `bitcoin rate infomation`,
	Run: func(cmd *cobra.Command, args []string) {
		errsCh := make(chan error)
		go func() {
			pid := os.Getpid()
			stdlog.Printf("start server!! this pid %d\n", pid)
			removePIDFile()
			savePID(pid)
			errsCh <- server.ApiNew(configName).Listen()
		}()
		stdlog.Fatalln("terminated", <-errsCh)
	},
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
