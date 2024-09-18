package services

import "github.com/danhbuidcn/go_backend_api/internal/repositories"

type UserService struct{
	userRepo *repositories.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepo(),
	}
}

func (us *UserService) GetinfoUserService() string {
	return us.userRepo.GetInfoUser()
}
