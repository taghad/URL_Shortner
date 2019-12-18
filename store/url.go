package store

import (
	"github.com/jinzhu/gorm"
	"github.com/taghad/URL_Shortner/model"
)

type UrlStore struct {
	db *gorm.DB
}
func NewUrlStore(db *gorm.DB) *UrlStore {
	return &UrlStore{
		db: db,
	}
}

func (u *UrlStore) Create(url *model.URL) (err error) {
	return u.db.Create(url).Error
}

func (u *UrlStore) Delete(url * model.URL) (err error)  {
	return u.db.Delete(u).Error
}

func (u *UrlStore) GetByShortURL(shortURL string) (*model.URL, error) {
	var m model.URL
	if err := u.db.Where(&model.URL{ShortURL: shortURL}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}