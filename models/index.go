package models

import (
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/memcache"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

const (
	psqlConn = "host=%v user=%v dbname=%v sslmode=disable password=%v port=%v search_path=public"
)

var db *gorm.DB
var sqlDB *sql.DB
var bm cache.Cache

func init() {
	err := connect()
	if err != nil {
		log.Fatal(err)
	}
	bm, err = cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		log.Fatal(err)
	}
}

func connect() (err error) {
	strConn := fmt.Sprintf(psqlConn, "db-staging.custodian.link", "orders", "orders", "orders", 5432)
	db, err = gorm.Open(postgres.Open(strConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err == nil {
		if sqlDB, err = db.DB(); err == nil {
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(10)
			sqlDB.SetConnMaxLifetime(5*time.Minute)
		}
	}
	return
}