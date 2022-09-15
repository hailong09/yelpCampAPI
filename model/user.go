package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" bson:"password,omitempty" validate: "required"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}


func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}