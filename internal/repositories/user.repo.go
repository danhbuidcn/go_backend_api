package repositories

import (
	"fmt"

	"github.com/danhbuidcn/go_backend_api/global"
	"github.com/danhbuidcn/go_backend_api/internal/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

func NewUserRepoSitory() IUserRepository {
	return &userRepository{}
}

// GetUserEmail implements IUserRepository
func (*userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '?' ORDER BY email
	row := global.Mdb.Table(TableNameGoCrmUser).Where("user_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	fmt.Println("Email row:", row)

	return row != NumberNull
}
