package r

import (
	"gorm.io/gorm"
	"url-shortner-service/entity"
)

type UrlRepository interface {
	Create(url *entity.Url) error

	Get(id uint64) (*entity.Url, error)

	Update(url *entity.Url) error

	Delete(id uint64) error

	GetAll() (*[]entity.Url, error)

	FindByCode(code string) (*entity.Url, error)
}

type UrlRepositoryImpl struct {
	db *gorm.DB
}

func (r UrlRepositoryImpl) Get(id uint64) (*entity.Url, error) {
	var url entity.Url
	tx := r.db.Where("id = ?", id).First(url)
	if tx.Error != nil {
		return &entity.Url{}, tx.Error
	}
	return &url, nil
}

func (r UrlRepositoryImpl) GetAll() (*[]entity.Url, error) {
	var urls []entity.Url
	tx := r.db.Find(urls)
	if tx.Error != nil {
		return &[]entity.Url{}, tx.Error
	}
	return &urls, nil
}

func (r UrlRepositoryImpl) Create(url *entity.Url) error {
	tx := r.db.Create(url)
	return tx.Error
}

func (r UrlRepositoryImpl) Update(url *entity.Url) error {
	tx := r.db.Updates(url)
	return tx.Error
}

func (r UrlRepositoryImpl) Delete(id uint64) error {
	tx := r.db.Unscoped().Delete(&entity.Url{}, id)
	return tx.Error
}

func (r UrlRepositoryImpl) FindByCode(code string) (*entity.Url, error) {
	var url entity.Url
	tx := r.db.Where("code = ?", code).First(&url)
	if tx.Error != nil {
		return &entity.Url{}, tx.Error
	}
	return &url, nil
}
