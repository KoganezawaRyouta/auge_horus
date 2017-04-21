package importer

// import (
// 	"time"
//
// 	httpClient "github.com/batch_coincheck/http/client"
// 	"github.com/batch_coincheck/orm"
// )
//
// // InsertExchangeRate db insert to orm.ExchangeRate
// func InsertExchangeRate(adapter *orm.GormAdapter, ctrades []httpClient.Ctrade) {
// 	for _, ctrade := range ctrades {
// 		trade := newTrade(ctrade)
// 		adapter.DB.Create(&trade)
// 		methodName(InsertTrade)
// 	}
// }
//
// func newExchangeRate(chTicker httpClient.Cticker) orm.Ticker {
// 	return orm.Ticker{
// 		Last:      chTicker.Last,
// 		Bid:       chTicker.Bid,
// 		Ask:       chTicker.Ask,
// 		High:      chTicker.High,
// 		Low:       chTicker.Low,
// 		Volume:    chTicker.Volume,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// }
