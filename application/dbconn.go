package application

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DbConfig struct {
	Name     string `json:"name"`
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

type DbConnection struct {
	Config     DbConfig
	Name       string
	Connection interface{}
}

func GetDBConnectionByName(connName string) *gorm.DB {

	val, ok := app.DbConnections[connName]
	if !ok {
		GetLogger().Fatal("connection is not defined, Connection Name: " + connName)
	}

	if val.Connection == nil {
		GetLogger().Fatal("connection error, Connection Name: " + connName)
	}

	return val.Connection.(*gorm.DB)
}

func setDbConnection(dbConfig DbConfig) (interface{}, error) {
	if dbConfig.Driver == "postgres" {
		return setPostgresDbConnection(dbConfig)
	}
	return nil, errors.New("Can't handle connection to driver " + dbConfig.Driver)
}

func setPostgresDbConnection(dbConfig DbConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Username,
		dbConfig.DbName, dbConfig.Password)

	app.logger.Info("Connecting to Postgres db at %", dsn)

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		app.logger.Fatalf("Can't connect to postgres db %s error %s", dsn, err)
		return nil, err
	}
	app.logger.Info("Connected to Postgres db successfully at %", dsn)

	return db, nil
}
