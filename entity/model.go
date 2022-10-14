package entity

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Url struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Code     string `json:"code" gorm:"unique; not null"`
	Clicks   uint64 `json:"clicks" gorm:"clicks"`
}

func Setup() {
	dsn := "host=localhost port=5432 user=test password=test dbname=url_shortner sslmode=disable"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Url{})
	if err != nil {
		panic(err)
	}
}
