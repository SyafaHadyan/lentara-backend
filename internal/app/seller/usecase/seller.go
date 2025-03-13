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

type SellerUseCaseItf interface {
	SellerRegister(register dto.SellerRegister) (dto.ResponseSellerRegister, error)
	SellerLogin(login dto.SellerLogin) (string, error)
	UpdateSellerInfo(seller dto.UpdateSellerInfo, sellerID uuid.UUID) (dto.ResponseUpdateSellerInfo, error)
	GetSellerInfo(sellerID uuid.UUID) (dto.GetSellerInfo, error)
	GetPublicSellerINfo(sellerID uuid.UUID) (dto.GetPublicSellerInfo, error)
}

type SellerUseCase struct {
	sellerRepo repository.SellerMySQLItf
	jwt        jwt.JWTItf
}

func NewSellerUseCase(sellerRepo repository.SellerMySQLItf, jwt *jwt.JWT) SellerUseCaseItf {
	return &SellerUseCase{
		sellerRepo: sellerRepo,
		jwt:        jwt,
	}
}

func (u SellerUseCase) SellerRegister(register dto.SellerRegister) (dto.ResponseSellerRegister, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.ResponseSellerRegister{}, fiber.NewError(http.StatusInternalServerError, "failed to hash seller password")
	}

	seller := entity.Seller{
		ID:             uuid.New(),
		Name:           register.Name,
		Email:          register.Email,
		Username:       register.Username,
		Password:       string(hashedPassword),
		StoreName:      register.StoreName,
		StoreLocation:  register.StoreLocation,
		PhoneNumber:    register.PhoneNumber,
		ProfilePicture: "https://static.vecteezy.com/system/resources/previews/026/619/142/original/default-avatar-profile-icon-of-social-media-user-photo-image-vector.jpg",
	}

	err = u.sellerRepo.SellerRegister(&seller)
	if err != nil {
		return dto.ResponseSellerRegister{}, fiber.NewError(http.StatusInternalServerError, "failed to create seller user")
	}

	return seller.ParseToDTOResponseSellerRegister(), nil
}

func (u SellerUseCase) SellerLogin(login dto.SellerLogin) (string, error) {
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

func (u SellerUseCase) UpdateSellerInfo(seller dto.UpdateSellerInfo, sellerID uuid.UUID) (dto.ResponseUpdateSellerInfo, error) {
	sellerUpdate := &entity.Seller{
		ID:             sellerID,
		Name:           seller.Name,
		Email:          seller.Email,
		Username:       seller.Username,
		Password:       seller.Password,
		StoreName:      seller.StoreName,
		StoreLocation:  seller.StoreLocation,
		PhoneNumber:    seller.PhoneNumber,
		ProfilePicture: seller.ProfilePicture,
	}

	err := u.sellerRepo.UpdateSellerInfo(sellerUpdate)
	if err != nil {
		return dto.ResponseUpdateSellerInfo{}, fiber.NewError(http.StatusInternalServerError, "failed to update seller info")
	}

	return sellerUpdate.ParseToDTOResponseUpdateSellerInfo(), nil
}

func (u SellerUseCase) GetSellerInfo(sellerID uuid.UUID) (dto.GetSellerInfo, error) {
	sellerInfo := entity.Seller{}

	err := u.sellerRepo.GetSellerInfo(&sellerInfo, sellerID)
	if err != nil {
		return dto.GetSellerInfo{}, fiber.NewError(http.StatusInternalServerError, "failed to get seller info")
	}

	return sellerInfo.ParseToDTOGetSellerInfo(), nil
}

func (u SellerUseCase) GetPublicSellerINfo(sellerID uuid.UUID) (dto.GetPublicSellerInfo, error) {
	sellerInfo := entity.Seller{}

	err := u.sellerRepo.GetSellerInfo(&sellerInfo, sellerID)
	if err != nil {
		return dto.GetPublicSellerInfo{}, fiber.NewError(http.StatusInternalServerError, "failed to get public seller info")
	}

	return sellerInfo.ParseToDTOGetPublicSellerInfo(), nil
}
