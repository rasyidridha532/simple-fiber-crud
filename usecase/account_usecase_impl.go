package usecase

import (
	"github.com/go-playground/validator/v10"
	"simple-fiber-crud/entity"
	"simple-fiber-crud/exception"
	"simple-fiber-crud/repository"
)

func NewAccountUseCase(accountRepository *repository.AccountRepository, validate *validator.Validate) AccountUseCase {
	return &accountUseCaseImpl{
		AccountRepository: *accountRepository,
		Validate:          validate,
	}
}

type accountUseCaseImpl struct {
	AccountRepository repository.AccountRepository
	Validate          *validator.Validate
}

func (a *accountUseCaseImpl) Create(request entity.AccountRequest) (response entity.Account) {
	err := a.Validate.Struct(request)
	exception.LogError(err)

	account := entity.Account{
		Id:          request.Id,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		AccountType: request.AccountType,
	}

	a.AccountRepository.Insert(account)

	response = entity.Account{
		Id:          account.Id,
		FirstName:   account.FirstName,
		LastName:    account.LastName,
		AccountType: account.AccountType,
	}

	return response
}

func (a *accountUseCaseImpl) List() (responses []entity.Account) {
	accounts := a.AccountRepository.GetAllUsers()
	for _, account := range accounts {
		responses = append(responses, entity.Account{
			Id:          account.Id,
			FirstName:   account.FirstName,
			LastName:    account.LastName,
			AccountType: account.AccountType,
		})
	}

	return responses
}

func (a *accountUseCaseImpl) User() (responses entity.Account) {
	account := a.AccountRepository.GetUser()

	return account
}

func (a *accountUseCaseImpl) DeleteUser(account entity.Account) {
	accountID := entity.Account{Id: account.Id}

	a.AccountRepository.DeleteUser(accountID)
}
