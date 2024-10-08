package repositories

import (
	"fmt"
	"time"

	"github.com/danhbuidcn/go_backend_api/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

// AddOTP implements IUserAuthRepository
func (u *userAuthRepository) AddOTP(email string, otp int, exirationTime int64) error {
	// panic("unimplemented")
	key := fmt.Sprintf("user:%s:otp", email) // user:email:otp
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(exirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
