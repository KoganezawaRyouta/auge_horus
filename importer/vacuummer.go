package importer

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"sync"
	"time"

	httpClient "github.com/KoganezawaRyouta/augehorus/http/client"
	"github.com/KoganezawaRyouta/augehorus/model"
	"github.com/KoganezawaRyouta/augehorus/orm"
)

type Vacuummer struct {
	dbAdapter *orm.GormAdapter
}

// NewVacuummer  init of Vacuummer
func NewVacuummer(dbAdapter *orm.GormAdapter) *Vacuummer {
	vacuummer := Vacuummer{}
	vacuummer.dbAdapter = dbAdapter
	return &vacuummer
}

// Vacuum it obtains the information of the trades and ticker from coincheck.jp,
// and register to DB
func (vc *Vacuummer) Run() {

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

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func methodName(i interface{}) string {
	iv := reflect.ValueOf(i)
	return runtime.FuncForPC(iv.Pointer()).Name()
}
