package repository

import (
	"gorm.io/gorm"
	"sync"
	"url-shortner-service/conf"
	"url-shortner-service/entity"
)

type UserRepository struct {
	db *gorm.DB
}

var singletonUserRepo *UserRepository
var userRepoLock = sync.RWMutex{}

func GetUserRepositoryInstance() (*UserRepository, error) {
	if singletonUserRepo == nil {
		userRepoLock.Lock()
		defer userRepoLock.Unlock()
		if singletonUserRepo == nil {
			db, err := conf.GetDb()
			if err != nil {
				return nil, err
			}
			singletonUserRepo = &UserRepository{db: db}
		}
	}
	return singletonUserRepo, nil
}

func (r *UserRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	tx := r.db.Where("username = ?", username).First(&user)
	return &user, tx.Error
}

func (r *UserRepository) Create(user *entity.User) error {
	tx := r.db.Create(user)
	return tx.Error
}
