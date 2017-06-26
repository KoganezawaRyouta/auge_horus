package config

import (
	"fmt"
	"os"
	"strconv"
)

const adapterNameMysql string = "mysql"
const adapterNameMysql2 string = "mysql2"

type DBConfig struct {
	Adapter  string
	Charset  string
	Encoding string
	Database string
	Username string
	Password string
	Host     string
	Port     int
	Pool     int
	LogFile  string
}

func DBConfigNew() *DBConfig {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("cannnot get db port value : " + err.Error())
	}
	pool, err := strconv.Atoi(os.Getenv("DB_POOL"))
	if err != nil {
		panic("cannnot get db pool value : " + err.Error())
	}
	return &DBConfig{
		Adapter:  os.Getenv("DB_ADAPTER"),
		Charset:  os.Getenv("DB_CHARSET"),
		Encoding: os.Getenv("DB_ENCORDING"),
		Database: os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USER_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Pool:     pool,
		LogFile:  os.Getenv("DB_LOGFILE"),
	}
}

// AdapterName get db adapter name
func (db *DBConfig) AdapterName() string {
	switch db.Adapter {
	case adapterNameMysql, adapterNameMysql2:
		return adapterNameMysql
	}
	return ""
}

// DSN get db dns
func (db *DBConfig) DSN() string {
	switch db.Adapter {
	case adapterNameMysql, adapterNameMysql2:
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			db.Username,
			db.Password,
			db.Host,
			db.Port,
			db.Database,
			db.Charset,
		)
	}
	return ""
}
