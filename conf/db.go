package conf

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"url-shortner-service/utils"
)

var db *gorm.DB
var mutex = sync.RWMutex{}

func GetDb() (*gorm.DB, error) { // singleton
	if db == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if db == nil {
			var err error
			dsn := utils.GetEnvOrDefault("DATABASE_URL", "host=localhost port=5432 user=test password=test dbname=url_shortner sslmode=disable")
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			return db, err
		}
	}
	return db, nil
}
