package orm

import (
	"fmt"
	"log"
	"os"

	"github.com/KoganezawaRyouta/augehorus/model"
	"github.com/KoganezawaRyouta/augehorus/settings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// GormAdapter adaptert of gorm.DB
type GormAdapter struct {
	DB     *gorm.DB
	Config *settings.Config
}

// NewGormAdapter  init of db adapter
func NewGormAdapter(configName string) *GormAdapter {
	adapter := GormAdapter{}
	adapter.LoadConfig(configName)
	adapter.InitDb()
	return &adapter
}

// InitDb init of gorm.DB
func (s *GormAdapter) InitDb() {
	db, err := gorm.Open(s.Config.DB.AdapterName(), s.Config.DB.DSN())
	if err != nil {
		panic(err.Error())
	}

	logfile, err := os.OpenFile("./tmp/development_db.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open development_db.log:" + err.Error())
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

// LoadConfig db settings
func (s *GormAdapter) LoadConfig(configName string) {
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("louding conf error: %s \n", err))
	}
	viper.Unmarshal(&s.Config)
}
