package store

import "github.com/jinzhu/gorm"
import "github.com/taghad/URL_Shortner/model"

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (u *UserStore) Create(us *model.User) (err error) {
	return u.db.Create(us).Error
}

func (u *UserStore) Delete(us * model.User) (err error)  {
	return u.db.Delete(us).Error
}

func (u *UserStore) GetByUsername(username string) (*model.User, error) {
	var m model.User
	if err := u.db.Where(&model.User{Username: username}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

