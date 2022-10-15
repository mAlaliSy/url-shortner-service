package entity

import "url-shortner-service/conf"

func MigrateDB() {
	var err error
	db, err := conf.GetDb()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Url{}, &User{})
	if err != nil {
		panic(err)
	}
}
