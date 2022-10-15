package repository

import (
	"gorm.io/gorm"
	"sync"
	"url-shortner-service/conf"
	"url-shortner-service/entity"
)

type UrlRepository struct {
	db *gorm.DB
}

var singletonUrlRepo *UrlRepository
var urlRepoLock = sync.RWMutex{}

func GetUrlRepositoryInstance() (*UrlRepository, error) {
	if singletonUrlRepo == nil {
		urlRepoLock.Lock()
		defer urlRepoLock.Unlock()
		if singletonUrlRepo == nil {
			db, err := conf.GetDb()
			if err != nil {
				return nil, err
			}
			singletonUrlRepo = &UrlRepository{db: db}
		}
	}
	return singletonUrlRepo, nil
}

func (r UrlRepository) Get(id uint64, userId uint64) (*entity.Url, error) {
	var url entity.Url
	tx := r.db.Where("id = ? and user_id = ?", id, userId).First(&url)
	if tx.Error != nil {
		return &entity.Url{}, tx.Error
	}
	return &url, nil
}

func (r UrlRepository) GetAllByUser(userId uint64) (*[]entity.Url, error) {
	var urls []entity.Url
	tx := r.db.Where("user_id = ?", userId).Find(&urls)
	if tx.Error != nil {
		return &[]entity.Url{}, tx.Error
	}
	return &urls, nil
}

func (r UrlRepository) Create(url *entity.Url) error {
	tx := r.db.Create(&url)
	return tx.Error
}

func (r UrlRepository) Update(url *entity.Url) error {
	tx := r.db.Updates(&url)
	return tx.Error
}

func (r UrlRepository) Delete(id uint64, userId uint64) error {
	tx := r.db.Unscoped().Where("user_id = ?", userId).Delete(&entity.Url{}, id)
	return tx.Error
}

func (r UrlRepository) FindByCode(code string) (*entity.Url, error) {
	var url entity.Url
	tx := r.db.Where("code = ?", code).First(&url)
	if tx.Error != nil {
		return &entity.Url{}, tx.Error
	}
	return &url, nil
}

func (r UrlRepository) IncrementClicks(id uint64) error {
	// execute increment in a single statement to avoid concurrency issues as multiple users may visit the url at the same time
	tx := r.db.Exec("UPDATE urls SET clicks = clicks + 1 WHERE id = ?", id)
	return tx.Error
}
