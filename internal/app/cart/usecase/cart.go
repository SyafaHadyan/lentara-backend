package usecase

import (
	"lentara-backend/internal/app/cart/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CartUsecaseItf interface {
	CreateCart(cart dto.CreateCart, userID uuid.UUID) (dto.CreateCart, error)
}

type CartUsecase struct {
	cartRepo repository.CartMySQLItf
}

func NewCartUsecase(cartRepo repository.CartMySQLItf) CartUsecaseItf {
	return &CartUsecase{
		cartRepo: cartRepo,
	}
}

func (u *CartUsecase) CreateCart(cart dto.CreateCart, userID uuid.UUID) (dto.CreateCart, error) {
	cartUser := entity.Cart{
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
	}

	err := u.cartRepo.CreateCart(&cartUser)
	if err != nil {
		return dto.CreateCart{}, fiber.NewError(http.StatusInternalServerError, "failed to create cart")
	}

	return cartUser.ParseToDTOCreateCart(), nil
}
