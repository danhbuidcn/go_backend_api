package repositories

import (
	"fmt"

	"github.com/danhbuidcn/go_backend_api/global"
	"github.com/danhbuidcn/go_backend_api/internal/database"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
	sqlc *database.Queries
}

func NewUserRepoSitory() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}

// GetUserEmail implements IUserRepository
func (up *userRepository) GetUserByEmail(email string) bool {
	fmt.Printf("GetUserByEmail: %s", email)
	user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}

	return user.UserID != NumberNull
}
