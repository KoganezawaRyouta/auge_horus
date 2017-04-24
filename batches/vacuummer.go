package batches

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	httpClient "github.com/KoganezawaRyouta/augehorus/http/client"
	"github.com/KoganezawaRyouta/augehorus/model"
	"github.com/KoganezawaRyouta/augehorus/orm"
	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/go-kit/kit/log"
)

type Vacuummer struct {
	dbAdapter *orm.GormAdapter
	config    *settings.Config
	logger    log.Logger
	elapsed   time.Duration
}

// NewVacuummer  init of Vacuummer
func NewVacuummer(config *settings.Config) *Vacuummer {
	vacuummer := Vacuummer{}
	rand.Seed(time.Now().UnixNano())
	vacuummer.elapsed = time.Since(time.Now())

	dbAdapter := orm.NewGormAdapter(config)
	vacuummer.dbAdapter = dbAdapter
	vacuummer.config = config

	logfile, err := os.OpenFile(vacuummer.config.Batch.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open " + vacuummer.config.Batch.LogFile + err.Error())
	}
	var logger log.Logger
	defer logfile.Close()
	logger = log.NewJSONLogger(log.NewSyncWriter(logfile))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	vacuummer.logger = logger
	return &vacuummer
}

// Vacuum it obtains the information of the trades and ticker from coincheck.jp,
// and register to DB
func (vc *Vacuummer) Run() {
	vc.logger.Log("importer start")

	wg := &sync.WaitGroup{}
	queue := make(chan func(*orm.GormAdapter) bool)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// generate two workers
	for i := 0; i < 2; i++ { // generate two workers
		wg.Add(1)
		go deQueue(ctx, wg, vc.dbAdapter, queue)
	}

	// enqueue
	enqueue(queue, importTicker)
	enqueue(queue, importTrades)

	wg.Wait()

	vc.logger.Log("elapsed: ", vc.elapsed)
	vc.logger.Log("importer end")
}

func enqueue(queue chan func(*orm.GormAdapter) bool, job func(*orm.GormAdapter) bool) {
	queue <- job
}

func deQueue(ctx context.Context, wg *sync.WaitGroup, dbAdapter *orm.GormAdapter, queue chan func(*orm.GormAdapter) bool) {
BREAK:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker exit")
			break BREAK
		case excuter := <-queue:
			t := excuter(dbAdapter)
			wg.Done()
			fmt.Println("execute: ", t)
		}
	}
}

func importTicker(dbAdapter *orm.GormAdapter) bool {
	cticker := httpClient.CoinCheckTicker()
	ticker := model.NewTickerForCticker(cticker)
	dbAdapter.DB.Create(&ticker)
	return true
}

func importTrades(dbAdapter *orm.GormAdapter) bool {
	ctrades := httpClient.CoinCheckTrades()
	for _, ctrade := range ctrades {
		trade := model.NewTradeForCtrade(ctrade)
		dbAdapter.DB.Create(&trade)
	}
	return true
}
