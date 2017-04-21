package cmd

import (
	"math/rand"
	"os"
	"time"

	"github.com/KoganezawaRyouta/augehorus/importer"
	"github.com/KoganezawaRyouta/augehorus/orm"
	"github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
)

var importerCmd = &cobra.Command{
	Use:   "importer",
	Short: "import bitcoin rate info to database",
	Long:  `import bitcoin rate info to database`,
	Run: func(cmd *cobra.Command, args []string) {
		logfile, err := os.OpenFile("./tmp/development_importer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic("cannnot open development_importer.log:" + err.Error())
		}
		defer logfile.Close()
		logger = log.NewJSONLogger(log.NewSyncWriter(logfile))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

		logger.Log("importer start")
		rand.Seed(time.Now().UnixNano())
		elapsed := time.Since(time.Now())
		dbAdapter := orm.NewGormAdapter(configName)

		importer.NewVacuummer(dbAdapter).Run()
		logger.Log("elapsed: ", elapsed)
		logger.Log("importer end")
	},
}
