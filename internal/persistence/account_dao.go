package persistence

import (
	"account-service/internal/model"
	"errors"

	"gorm.io/gorm"
)

type AccountDAO interface {
	Create(account model.Account) (model.Account, error)
	Update(account model.Account) (model.Account, error)
	Get(id int) (model.Account, error)
	Delete(id int) error
	GetByUserId(userId int, page int, size int) (Page[model.Account], error)
}

type AccountDAOImpl struct {
	db *gorm.DB
}

func NewAccountDAO(conn *gorm.DB) AccountDAO {
	return AccountDAOImpl{conn}
}

func (dao AccountDAOImpl) Get(id int) (model.Account, error) {
	var account model.Account
	tx := dao.db.First(&account, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.Account{}, RecordNotFoundError{}
		}
		return model.Account{}, tx.Error
	}
	return account, nil
}

func (dao AccountDAOImpl) Create(account model.Account) (model.Account, error) {
	tx := dao.db.Create(&account)
	if tx.Error != nil {
		return model.Account{}, tx.Error
	}
	return account, nil
}

func (dao AccountDAOImpl) Update(account model.Account) (model.Account, error) {
	tx := dao.db.Save(&account)
	if tx.Error != nil {
		return model.Account{}, tx.Error
	}
	return account, nil
}

func (dao AccountDAOImpl) Delete(id int) error {
	tx := dao.db.Delete(&model.Account{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (dao AccountDAOImpl) GetByUserId(userId int, page int, size int) (Page[model.Account], error) {
	var accounts []model.Account = make([]model.Account, 1)
	tx := dao.db.Scopes(Paginate(page, size)).Find(&accounts, "user_id = ?", userId)
	if tx.Error != nil {
		return Page[model.Account]{}, tx.Error
	}
	var total int64
	dao.db.Model(model.Account{}).Where("user_id = ?", userId).Count(&total)
	return Page[model.Account]{
		List:  accounts,
		Page:  page,
		Size:  len(accounts),
		Total: total,
	}, nil
}
