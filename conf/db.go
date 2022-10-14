package conf

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB
var mutex = sync.RWMutex{}

func initDb() {
	var err error
	dsn := "host=localhost port=5432 user=test password=test dbname=url_shortner sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDb() *gorm.DB { // singleton
	if db == nil {
		mutex.Lock()
		if db == nil {
			initDb()
		}
		mutex.Unlock()
	}
	return db
}
