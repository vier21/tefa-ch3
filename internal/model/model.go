package model

type User struct {
	UserID  string `db:"id" json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `db:"name" json:"name" bson:"name"`
	Address string `db:"address" json:"address" bson:"address"`
	Email   string `db:"email" json:"email" bson:"email"`
}
