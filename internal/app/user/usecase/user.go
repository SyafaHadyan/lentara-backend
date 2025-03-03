package usecase

import (
	"lentara-backend/internal/app/user/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseItf interface {
	Register(dto.Register) (dto.ResponseRegister, error)
}

type UserUsecase struct {
	userRepo repository.UserMySQLItf
}

func NewUserUsecase(userRepo repository.UserMySQLItf) UserUsecaseItf {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Register(register dto.Register) (dto.ResponseRegister, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.ResponseRegister{}, fiber.NewError(http.StatusInternalServerError, "failed to hash user password")
	}

	user := entity.User{
		ID:       uuid.New(),
		Name:     register.Name,
		Email:    register.Email,
		Password: string(hashedPassword),
		IsAdmin:  register.IsAdmin,
	}

	err = u.userRepo.Create(&user)
	if err != nil {
		return dto.ResponseRegister{}, fiber.NewError(http.StatusInternalServerError, "failed to create user")
	}

	return user.ParseToDTOResponseRegister(), nil
}
