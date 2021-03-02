package model

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	. "kamgo/config"
	"log"
	"os"
	"time"
)

var (
	db       *gorm.DB
	dbLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)
)

func KamDB() *gorm.DB {
	if db == nil {
		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.AutoMigrate(KamSubscriber{})
		newDb.Logger = newDb.Logger.LogMode(logger.Info)
		db = newDb
	}
	return db
}

func newDB() (*gorm.DB, error) {
	sqlConnection := Conf.KamDB.UserName + ":" + Conf.KamDB.Pwd + "@tcp(" + Conf.KamDB.Host + ":" + Conf.KamDB.Port + ")/" + Conf.KamDB.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	if sqlDB, err := sql.Open("mysql", sqlConnection); err != nil {
		return nil, err
	} else {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		if gormDB, err := gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{Logger: dbLogger}); err != nil {
			return nil, err
		} else {
			return gormDB, nil
		}
	}
}
