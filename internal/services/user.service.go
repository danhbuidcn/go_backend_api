package services

import (
	"github.com/danhbuidcn/go_backend_api/internal/repositories"
	"github.com/danhbuidcn/go_backend_api/pkg/response"
)

// type UserService struct{
// 	userRepo *repositories.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repositories.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetinfoUserService() string {
// 	return us.userRepo.GetInfoUser()
// }

// INTERFACE VERSION

type IUserService interface {
	Resgister(email string, purpose string) int
}

type userService struct {
	userRepository repositories.IUserRepository
	// declare all interfaces here
}

// Resgister implements IUserService.
func (us userService) Resgister(email string, purpose string) int {
	panic("unimplemented")
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

// Register implements IUserService
func (us *userService) Register(email string, purpose string) int {
	// check email exists
	if us.userRepository.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrorCodeSuccess
}
