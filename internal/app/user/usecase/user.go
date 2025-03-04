package usecase

import (
	"lentara-backend/internal/app/user/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"lentara-backend/internal/infra/jwt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseItf interface {
	Register(dto.Register) (dto.ResponseRegister, error)
	Login(dto.Login) (string, error)
}

type UserUsecase struct {
	userRepo repository.UserMySQLItf
	jwt      jwt.JWTItf
}

func NewUserUsecase(userRepo repository.UserMySQLItf, jwt *jwt.JWT) UserUsecaseItf {
	return &UserUsecase{
		userRepo: userRepo,
		jwt:      jwt,
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
		Username: register.Username,
		Password: string(hashedPassword),
		IsAdmin:  false,
	}

	err = u.userRepo.Create(&user)
	if err != nil {
		return dto.ResponseRegister{}, fiber.NewError(http.StatusInternalServerError, "failed to create user")
	}

	return user.ParseToDTOResponseRegister(), nil
}

func (u *UserUsecase) Login(login dto.Login) (string, error) {
	var user entity.User

	err := u.userRepo.Get(&user, dto.UserParam{Username: login.Username})
	if err != nil {
		return "", fiber.NewError(http.StatusBadRequest, "username or password is invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", fiber.NewError(http.StatusBadRequest, "username or password is invalid")
	}

	token, err := u.jwt.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return "", fiber.NewError(http.StatusInternalServerError, "failed to generate token")
	}

	return token, nil
}

// func (u *UserUsecase) Login(login dto.Login) (dto.Reesponselogin, error) {
// 	err := bcrypt.CompareHashAndPassword(repository.UserMySQL, []byte(login.Password))
// 	if err != nil {
// 		return dto.ResponseLogin{}, fiber.NewError(http.StatusBadRequest, "invalid password")
// 	}
//
// 	config, err := env.New()
// 	if err != nil {
// 		return dto.ResponseLogin{}, fiber.NewError(http.StatusInternalServerError, "failed to get env")
// 	}
//
// 	token = jwt.New(jwt.SigningMethodHS256)
// 	s = token.SignedString(config.JWTSecretKey)
// }
