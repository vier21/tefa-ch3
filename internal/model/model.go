package model

type User struct {
	UserID  string `db:"id" json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `db:"name" json:"name" bson:"name"`
	Address string `db:"address" json:"address" bson:"address"`
	Email   string `db:"email" json:"email" bson:"email"`
}

type Account struct {
	AccountID      string `db:"id" json:"id,omitempty" bson:"_id,omitempty"`
	MsisdnCustomer string `db:"msisdn_customer" json:"msisdn_customer" bson:"msisdn_customer"`
	UserID         string `db:"user_id" json:"user_id" bson:"user_id"`
}
