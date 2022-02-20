package repository

import "simple-fiber-crud/entity"

type AccountRepository interface {
	GetAllUsers() (account []entity.Account)
	GetUser() (account entity.Account)
	Insert(account entity.Account)
	DeleteUser(delete entity.Account)
}
