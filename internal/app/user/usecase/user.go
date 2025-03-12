package usecase

import (
	"lentara-backend/internal/app/user/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"lentara-backend/internal/infra/jwt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseItf interface {
	Register(dto.Register) (dto.ResponseRegister, error)
	Login(dto.Login) (string, error)
	GetUserInfoByUserID(userID uuid.UUID) (dto.GetUserInfoByUserID, error)
	UpdateUserInfo(user dto.UpdateUserInfo, userID uuid.UUID) (dto.UpdateUserInfo, error)
}

type UserUseCase struct {
	userRepo repository.UserMySQLItf
	jwt      jwt.JWTItf
}

func NewUserUseCase(userRepo repository.UserMySQLItf, jwt *jwt.JWT) UserUseCaseItf {
	return &UserUseCase{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (u UserUseCase) Register(register dto.Register) (dto.ResponseRegister, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.ResponseRegister{}, fiber.NewError(http.StatusInternalServerError, "failed to hash user password")
	}

	user := entity.User{
		ID:             uuid.New(),
		Name:           register.Name,
		Email:          register.Email,
		Username:       register.Username,
		Password:       string(hashedPassword),
		IsAdmin:        false,
		ProfilePicture: "https://static.vecteezy.com/system/resources/previews/026/619/142/original/default-avatar-profile-icon-of-social-media-user-photo-image-vector.jpg",
	}

	err = u.userRepo.RegisterUser(&user)
	if err != nil {
		return dto.ResponseRegister{}, fiber.NewError(http.StatusInternalServerError, "failed to create user")
	}

	return user.ParseToDTOResponseRegister(), nil
}

func (u UserUseCase) Login(login dto.Login) (string, error) {
	var user entity.User

	err := u.userRepo.GetUserUsername(&user, dto.UserParam{Username: login.Username})
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

func (u UserUseCase) UpdateUserInfo(user dto.UpdateUserInfo, userID uuid.UUID) (dto.UpdateUserInfo, error) {
	userUpdate := entity.User{
		ID:           userID,
		Name:         user.Name,
		Email:        user.Email,
		Username:     user.Username,
		Password:     user.Password,
		UserLocation: user.UserLocation,
		PhoneNumber:  user.PhoneNumber,
	}

	err := u.userRepo.UpdateUserInfo(&userUpdate)
	if err != nil {
		return dto.UpdateUserInfo{}, fiber.NewError(http.StatusInternalServerError, "failed to update user info")
	}

	err = u.userRepo.GetUserInfoByUserID(&userUpdate, userID)
	if err != nil {
		// Continue even if failed to get updated user data
		log.Println("failed to get updated user data")
	}

	return userUpdate.ParseToDTOUpdateUserInfo(), nil
}

func (u UserUseCase) GetUserInfoByUserID(userID uuid.UUID) (dto.GetUserInfoByUserID, error) {
	user := entity.User{
		ID: userID,
	}

	err := u.userRepo.GetUserInfoByUserID(&user, userID)
	if err != nil {
		return dto.GetUserInfoByUserID{}, fiber.NewError(http.StatusInternalServerError, "failed to get user info by user id")
	}

	return user.ParseToDTOGetUserInfoByUserID(), nil
}
