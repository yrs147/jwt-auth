package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID			primitive		`bson: "_id"`
	username	*string			`json:"first_name" validate:"re"`
	Password	
	Email
	Token
	User_type
	Refresh_token
	Created_at
	Updated_at
	User_id
}