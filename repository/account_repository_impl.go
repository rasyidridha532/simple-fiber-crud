package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-fiber-crud/config"
	"simple-fiber-crud/entity"
	"simple-fiber-crud/exception"
)

var logging = config.Logger()

func NewAccountRepository(database *mongo.Database) AccountRepository {
	return &accountRepositoryImpl{
		Collection: database.Collection("account"),
	}
}

type accountRepositoryImpl struct {
	Collection *mongo.Collection
}

func (a *accountRepositoryImpl) GetAllUsers() (account []entity.Account) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := a.Collection.Find(ctx, bson.M{})
	exception.LogError(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.LogError(err)

	for _, document := range documents {
		account = append(account, entity.Account{
			Id:          document["_id"].(string),
			FirstName:   document["first_name"].(string),
			LastName:    document["last_name"].(string),
			AccountType: document["account_type"].(string),
		})
	}

	return account
}

func (a *accountRepositoryImpl) GetUser() (account entity.Account) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor := a.Collection.FindOne(ctx, bson.M{"_id": account.Id})
	err := cursor.Decode(&account)
	exception.LogError(err)

	return account
}

func (a *accountRepositoryImpl) Insert(account entity.Account) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := a.Collection.InsertOne(ctx, bson.M{
		"_id":          account.Id,
		"first_name":   account.FirstName,
		"last_name":    account.LastName,
		"account_type": account.AccountType,
	})
	exception.LogError(err)

	logging.Info("Insert Success!")
	defer func() {
		err = logging.Sync()
		exception.LogError(err)
	}()
}

func (a *accountRepositoryImpl) DeleteUser(account entity.Account) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := a.Collection.DeleteOne(ctx, bson.M{"_id": account.Id})
	exception.LogError(err)

	logging.Info("User has been deleted!")
	defer func() {
		err = logging.Sync()
		exception.LogError(err)
	}()
}
