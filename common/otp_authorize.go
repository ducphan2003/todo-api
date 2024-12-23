package common

import (
	"time"
	"todo-api/component/hasher"
)

func OTPAuthorize(verificationCode string, userSalt string, userVerificationCode string, createdVerificationCodeAt *time.Time) error {
	timeNow := time.Now()
	md5 := hasher.NewMd5Hash()
	VerificationCodeHashed := md5.Hash(verificationCode + userSalt)
	if VerificationCodeHashed != userVerificationCode {
		return NewCustomError(nil, "OTP is not correct", "OTPIsNotCorrect")
	}
	timeMinus := timeNow.Sub(*createdVerificationCodeAt)
	if timeMinus.Seconds() > 300 {
		return NewCustomError(nil, "OTP is expired", "OTPIsExpired")
	}
	return nil
}
