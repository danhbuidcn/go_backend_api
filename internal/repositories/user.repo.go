package repositories

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "string"
// }

// INTERFACE VERSION

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

func NewUserRepoSitory() IUserRepository {
	return &userRepository{}
}

// GetUserEmail implements IUserRepository
func (*userRepository) GetUserByEmail(email string) bool {
	panic("unimplement")
}
