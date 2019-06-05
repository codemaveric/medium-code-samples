package contact

import "time"

type Contact struct {
	FirstName   string    `json:"firstName" bson:"firstName"`
	LastName    string    `json:"lastName" bson:"lastName"`
	Email       string    `json:"email" bson:"email"`
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	Address     string    `json:"address" bson:"address"`
	Company     string    `json:"company" bson:"company"`
	CreatedOn   time.Time `json:"createdOn" bson:"createdon"`
}
