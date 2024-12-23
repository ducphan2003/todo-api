package taskmodel

import (
	"errors"
	"todo-api/database"
	userModel "todo-api/modules/user/usermodel"
)

const TaskEntityName = "Tasks"

type Task struct {
	database.BaseModel `json:",inline"`
	UserID             uint            `json:"user_id" gorm:"column:user_id"`
	User               *userModel.User `json:"user" gorm:"foreignKey:UserID;references:ID;"`
	Title              string          `json:"title" gorm:"title"`
	Description        string          `json:"description" gorm:"description"`
	Progress           string          `json:"progress" gorm:"column:progress;"`
	Priority           string          `json:"priority" gorm:"column:priority;"`
}

func (Task) TableName() string {
	return "tasks"
}

type TaskCreate struct {
	database.BaseModel `json:",inline"`
	UserID             uint   `json:"user_id" gorm:"column:user_id"`
	Title              string `json:"title" gorm:"column:title"`
	Description        string `json:"description" gorm:"column:description"`
	Progress           string `json:"progress" gorm:"column:progress;"`
	Priority           string `json:"priority" gorm:"column:priority"`
}

type TaskUpdate struct {
	database.BaseModel `json:",inline"`
	UserID             uint   `json:"user_id" gorm:"column:user_id"`
	Title              string `json:"title" gorm:"column:title"`
	Description        string `json:"description" gorm:"column:description"`
	Progress           string `json:"progress" gorm:"column:progress;"`
	Priority           string `json:"priority" gorm:"column:priority"`
}

func (TaskUpdate) TableName() string {
	return "tasks"
}

type TaskFilter struct {
	Status      string `json:"status" form:"status"`
	UserID      uint   `json:"user_id" form:"user_id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Progress    string `json:"progress" form:"progress;"`
	Priority    string `json:"priority" form:"priority"`
}

func (u Task) Validate() error {
	if u.Title == "" {
		return errors.New("title is required")
	}
	return nil
}
