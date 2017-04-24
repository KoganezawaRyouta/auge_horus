package orm

import (
	"log"
	"os"

	"github.com/KoganezawaRyouta/augehorus/model"
	"github.com/KoganezawaRyouta/augehorus/settings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// GormAdapter adaptert of gorm.DB
type GormAdapter struct {
	DB     *gorm.DB
	Config *settings.Config
}

// NewGormAdapter  init of db adapter
func NewGormAdapter(config *settings.Config) *GormAdapter {
	adapter := GormAdapter{}
	adapter.Config = config
	adapter.InitDb()
	return &adapter
}

// InitDb init of gorm.DB
func (s *GormAdapter) InitDb() {
	db, err := gorm.Open(s.Config.DB.AdapterName(), s.Config.DB.DSN())
	if err != nil {
		panic(err.Error())
	}

	logfile, err := os.OpenFile(s.Config.DB.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open " + s.Config.DB.LogFile + err.Error())
	}
	db.SetLogger(log.New(logfile, "****", log.LstdFlags|log.Llongfile))
	s.DB = db
	s.DB.LogMode(true)
}

// InitSchema db migrate
func (s *GormAdapter) InitSchema() {
	s.DB.AutoMigrate(&model.Ticker{})
	s.DB.AutoMigrate(&model.Trade{})
}
