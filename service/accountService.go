package service

import (
	"banking-app/domain"
	"banking-app/dto"
	"banking-app/errs"
)

type AccountService interface {
	Save(dto.AccountRequest) (dto.AccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) Save(account dto.AccountRequest) (dto.AccountResponse, *errs.AppError) {

	entity := domain.Account{CustomerId: account.CustomerId, OpeningDate: account.OpeningDate, AccountType: account.AccountType, Amount: account.Amount, Status: account.StatusToString()}

	savedEntity, _ := s.repo.Save(entity)

	return savedEntity.ToDto(), nil

}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository}
}
