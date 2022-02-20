package entity

type Account struct {
	Id          string `json:"id,omitempty" validate:"required" bson:"_id"`
	FirstName   string `json:"first_name" validate:"required" bson:"first_name"`
	LastName    string `json:"last_name" validate:"required" bson:"last_name"`
	AccountType string `json:"account_type" validate:"required" bson:"account_type"`
}

type AccountRequest struct {
	Id          string `json:"id,omitempty" validate:"required" bson:"_id"`
	FirstName   string `json:"first_name,omitempty" validate:"required" bson:"first_name"`
	LastName    string `json:"last_name,omitempty" validate:"required" bson:"last_name"`
	AccountType string `json:"account_type,omitempty" validate:"required" bson:"account_type"`
}
