package usecase

import "simple-fiber-crud/entity"

type AccountUseCase interface {
	Create(request entity.AccountRequest) (response entity.Account)
	List() (responses []entity.Account)
	User() (responses entity.Account)
	DeleteUser(account entity.Account)
}
