package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID				primitive		`bson: "_id"`
	Username		*string			`json:"username" validate:"required, min=2, max=100"`
	Password		*string			`json:"Password" validate:"required, min=6"`
	Email			*string			`json:"email" validate:"email, required"`
	Token			*string			`json: "token"`
	User_type		*string			`json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Refresh_token	*string			`json:"refresh_token"`
	Created_at		time.Time 		`json:"created_at"`
	Updated_at		time.Time		`json:"update_at"`
	User_id			string			`json:"user_id"`
}