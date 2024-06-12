package service

import (
	"account-service/internal/model"
	"account-service/internal/persistence"
	"errors"
)

type AccountService interface {
	GetById(userId int64, id int64) (model.Account, error)
	GetAll(userId int64, page int, size int) (persistence.Page[model.Account], error)
	Create(userId int64, account model.Account) (model.Account, error)
	Update(userId int64, id int64, account model.Account) (model.Account, error)
	Delete(userId int64, id int64) error
}

type AccountServiceImpl struct {
	dao persistence.AccountDAO
}

func NewAccountService(dao persistence.AccountDAO) AccountService {
	return AccountServiceImpl{dao}
}

func (s AccountServiceImpl) GetById(userId int64, id int64) (model.Account, error) {
	account, err := s.dao.Get(id)
	if err != nil {
		if errors.Is(err, persistence.RecordNotFoundError{}) {
			return model.Account{}, NotFoundError{Type: model.Account{}, Id: id}
		}
		return model.Account{}, err
	}
	if account.UserId != userId {
		return model.Account{}, NotFoundError{Type: model.Account{}, Id: id}
	}
	return account, nil
}

func (s AccountServiceImpl) GetAll(userId int64, page int, size int) (persistence.Page[model.Account], error) {
	return s.dao.GetByUserId(userId, page, size)
}

func (s AccountServiceImpl) Create(userId int64, account model.Account) (model.Account, error) {
	account.ID = 0
	account.UserId = userId
	return s.dao.Create(account)
}

func (s AccountServiceImpl) Update(userId int64, id int64, account model.Account) (model.Account, error) {
	found, err := s.GetById(userId, id)
	if err != nil {
		return model.Account{}, err
	}
	account.ID = id
	account.UserId = userId
	account.CreatedAt = found.CreatedAt
	return s.dao.Update(account)
}

func (s AccountServiceImpl) Delete(userId int64, id int64) error {
	_, err := s.GetById(userId, id)
	if err != nil {
		return err
	}
	return s.dao.Delete(id)
}
