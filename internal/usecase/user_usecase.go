package usecase

import (
	// "errors"
	
	"nutrimama/internal/entity"
	"nutrimama/internal/model"
	"nutrimama/internal/model/converter"
	"nutrimama/internal/repository"
	"nutrimama/internal/utils"

	bcrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB              *gorm.DB
	UserRepository  *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, UserRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		UserRepository: UserRepository,
	}
}

func (u *UserUseCase) Create(req model.RegisterRequest) (*model.UserResponse, error) {
	c := u.DB.Begin()
	// existingUser, _ := u.UserRepository.FindByEmail(req.Email)
	// if existingUser != nil {
	// 	return nil, errors.New("email already exists")
	// }
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashedPassword)
	user := entity.User{
		Email:      req.Email,
		Password:   req.Password,
		PhoneNumber: req.PhoneNumber,
		Role:      req.Role,
	}

	if err := u.UserRepository.Create(c, &user); err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return converter.UserToResponse(&user, token), nil

}