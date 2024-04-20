package service

import (
	"account-service/internal/model"
	"account-service/internal/persistence"
	"errors"
)

type AccountService interface {
	GetById(userId int, id int) (model.Account, error)
	GetAll(userId int, page int, size int) (persistence.Page[model.Account], error)
	Create(userId int, account model.Account) (model.Account, error)
	Update(userId int, id int, account model.Account) (model.Account, error)
	Delete(userId int, id int) error
}

type AccountServiceImpl struct {
	dao persistence.AccountDAO
}

func NewAccountService(dao persistence.AccountDAO) AccountService {
	return AccountServiceImpl{dao}
}

func (s AccountServiceImpl) GetById(userId int, id int) (model.Account, error) {
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

func (s AccountServiceImpl) GetAll(userId int, page int, size int) (persistence.Page[model.Account], error) {
	return s.dao.GetByUserId(userId, page, size)
}

func (s AccountServiceImpl) Create(userId int, account model.Account) (model.Account, error) {
	account.ID = 0
	account.UserId = userId
	return s.dao.Create(account)
}

func (s AccountServiceImpl) Update(userId int, id int, account model.Account) (model.Account, error) {
	_, err := s.GetById(userId, id)
	if err != nil {
		return model.Account{}, err
	}
	account.ID = id
	account.UserId = userId
	return s.dao.Update(account)
}

func (s AccountServiceImpl) Delete(userId int, id int) error {
	_, err := s.GetById(userId, id)
	if err != nil {
		return err
	}
	return s.dao.Delete(id)
}
