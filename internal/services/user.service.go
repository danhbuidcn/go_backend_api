package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/danhbuidcn/go_backend_api/internal/repositories"
	"github.com/danhbuidcn/go_backend_api/internal/utils/crypto"
	"github.com/danhbuidcn/go_backend_api/internal/utils/random"
	"github.com/danhbuidcn/go_backend_api/internal/utils/sendto"
	"github.com/danhbuidcn/go_backend_api/pkg/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepository repositories.IUserRepository
	userAuthRepo   repositories.IUserAuthRepository
	// declare all interfaces here
}

func NewUserService(
	userRepository repositories.IUserRepository,
	userAuthRepo repositories.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepository: userRepository,
		userAuthRepo:   userAuthRepo,
	}
}

// Register implements IUserService
func (us *userService) Register(email string, purpose string) int {
	// 0.hash email
	hashEmail := crypto.GetHash(email)
	fmt.Sprintf("hashEmail::%s", hashEmail)

	// 5.check OTP is available

	// 6.handle user spam

	// 1.check email exists in db
	if us.userRepository.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2.new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("Otp is ::: %d\n", otp)

	// 3.save OTP in redis with exiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}

	// // 4.send email OTP
	// // make sure the email sent is valid
	// err = sendto.SendTextEmailOtp([]string{email}, os.Getenv("SENDER_EMAIL"), strconv.Itoa(otp))
	// if err != nil {
	// 	return response.ErrSendEmailOtp
	// }

	// Or 4.send OTP by JAVA
	err = sendto.SendEmailToJavaByAPI(strconv.Itoa(otp), email, "otp-auth.html")
	if err != nil {
		return response.ErrSendEmailOtp
	}

	return response.ErrorCodeSuccess
}
