package repository

import (
	"gorm.io/gorm"
	"sync"
	"url-shortner-service/conf"
	"url-shortner-service/entity"
)

type UserRepository interface {
	FindByUsername(username string) (*entity.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

var singletonUserRepo *UserRepositoryImpl
var userRepoLock = sync.RWMutex{}

func GetUserRepositoryInstance() (*UserRepositoryImpl, error) {
	if singletonUserRepo == nil {
		userRepoLock.Lock()
		defer userRepoLock.Unlock()
		if singletonUserRepo == nil {
			db, err := conf.GetDb()
			if err != nil {
				return nil, err
			}
			singletonUserRepo = &UserRepositoryImpl{db: db}
		}
	}
	return singletonUserRepo, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	tx := r.db.Where("username = ?", username).First(&user)
	return &user, tx.Error
}
