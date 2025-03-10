package usecase

import (
	"lentara-backend/internal/app/seller/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"lentara-backend/internal/infra/jwt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SellerUsecaseItf interface {
	SellerRegister(register dto.SellerRegister) (dto.ResponseSellerRegister, error)
	SellerLogin(login dto.SellerLogin) (string, error)
	UpdateSellerInfo(seller dto.UpdateSellerInfo, sellerID uuid.UUID) (dto.UpdateSellerInfo, error)
	GetSellerInfo(seller dto.GetSellerInfo, sellerID uuid.UUID) (dto.GetSellerInfo, error)
}

type SellerUsecase struct {
	sellerRepo repository.SellerMySQLItf
	jwt        jwt.JWTItf
}

func NewSellerUsecase(sellerRepo repository.SellerMySQLItf, jwt *jwt.JWT) SellerUsecaseItf {
	return &SellerUsecase{
		sellerRepo: sellerRepo,
		jwt:        jwt,
	}
}

func (u *SellerUsecase) SellerRegister(register dto.SellerRegister) (dto.ResponseSellerRegister, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.ResponseSellerRegister{}, fiber.NewError(http.StatusInternalServerError, "failed to hash seller password")
	}

	seller := entity.Seller{
		ID:       uuid.New(),
		Name:     register.Name,
		Email:    register.Email,
		Username: register.Username,
		Password: string(hashedPassword),
	}

	err = u.sellerRepo.SellerRegister(&seller)
	if err != nil {
		return dto.ResponseSellerRegister{}, fiber.NewError(http.StatusInternalServerError, "failed to create seller user")
	}

	return seller.ParseToDTOResponseSellerRegister(), nil
}

func (u *SellerUsecase) SellerLogin(login dto.SellerLogin) (string, error) {
	var seller entity.Seller

	err := u.sellerRepo.GetSellerLoginInfo(&seller, dto.SellerInfo{Username: login.Username})
	if err != nil {
		return "", fiber.NewError(http.StatusBadRequest, "username or password is invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(seller.Password), []byte(login.Password))
	if err != nil {
		return "", fiber.NewError(http.StatusBadRequest, "username or password is invalid")
	}

	token, err := u.jwt.GenerateToken(seller.ID, false)
	if err != nil {
		return "", fiber.NewError(http.StatusInternalServerError, "failed to generate token")
	}

	return token, nil
}

func (u *SellerUsecase) UpdateSellerInfo(seller dto.UpdateSellerInfo, sellerID uuid.UUID) (dto.UpdateSellerInfo, error) {
	sellerUpdate := &entity.Seller{
		ID:            sellerID,
		Name:          seller.Name,
		Email:         seller.Email,
		Username:      seller.Username,
		Password:      seller.Password,
		StoreLocation: seller.StoreLocation,
		PhoneNumber:   seller.PhoneNumber,
	}

	err := u.sellerRepo.UpdateSellerInfo(sellerUpdate)
	if err != nil {
		return dto.UpdateSellerInfo{}, fiber.NewError(http.StatusInternalServerError, "failed to update seller info")
	}

	return sellerUpdate.ParseToDTOResponseUpdateSellerInfo(), nil
}

func (u *SellerUsecase) GetSellerInfo(seller dto.GetSellerInfo, sellerID uuid.UUID) (dto.GetSellerInfo, error) {
	sellerInfo := &entity.Seller{
		ID:            sellerID,
		Name:          seller.Name,
		Email:         seller.Email,
		Username:      seller.Username,
		Password:      seller.Password,
		StoreLocation: seller.StoreLocation,
		PhoneNumber:   seller.PhoneNumber,
		CreatedAt:     seller.CreatedAt,
		UpdatedAt:     seller.UpdatedAt,
	}

	err := u.sellerRepo.GetSellerInfo(sellerInfo)
	if err != nil {
		return dto.GetSellerInfo{}, fiber.NewError(http.StatusInternalServerError, "failed to get seller info")
	}

	return sellerInfo.ParseToDTOGetSellerInfo(), nil
}
