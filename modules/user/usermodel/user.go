package usermodel

import (
	"errors"
	"todo-api/database"
)

const UserEntityName = "Users"

type User struct {
	database.BaseModel `json:",inline"`
	Name               string `json:"name" gorm:"column:name,uniqueIndex"`
	Password           string `json:"-" gorm:"column:password;"`
	Salt               string `json:"-" gorm:"column:salt;"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	database.BaseModel `json:",inline"`
	Name               string `json:"name" gorm:"column:name"`
	Password           string `gorm:"column:password"`
	Salt               string `gorm:"column:salt"`
}

type UserUpdate struct {
	database.BaseModel `json:",inline"`
	Name               string `json:"name" gorm:"column:name"`
	Password           string `json:"password" gorm:"column:password"`
}

type UserFilter struct {
	database.BaseModel `json:",inline"`
	Name               string `json:"name" form:"name"`
}

func (u User) Validate() error {

	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
