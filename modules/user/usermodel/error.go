package usermodel

import (
	"errors"
	"todo-api/common"
)

var (
	ErrUserNameOrPasswordInvalid = common.NewCustomError(
		errors.New("user name or password is invalid"),
		"username or password is invalid",
		"ErrUserNameOrPasswordInvalid",
	)
	PasswordNotCorrect = common.NewCustomError(
		errors.New("password is not correct"),
		"password is not correct",
		"ErrPasswordNotCorrect",
	)
	VerificationCodeIncorrectOrExpired = common.NewCustomError(
		errors.New("verification code is incorrect or expired"),
		"verification code is incorrect",
		"ErrVerificationCodeIncorrectOrExpired",
	)
	AccountDeletedOrNotFound = common.NewCustomError(
		errors.New("account is deleted or not found"),
		"account is deleted or not found",
		"ErrAccountDeletedOrNotFound",
	)
	AccountDeleted = common.NewCustomError(
		errors.New("account is deleted"),
		"account is deleted",
		"ErrAccountDeleted",
	)
	AccountNotFound = common.NewCustomError(
		errors.New("account is not found"),
		"account is not found",
		"ErrAccountNotFound",
	)
)
