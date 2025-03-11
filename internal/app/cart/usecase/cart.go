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
	CreateCart(cart dto.CreateCart, userID uuid.UUID, sellerID uuid.UUID) (dto.CreateCart, error)
	UpdateCart(cart dto.UpdateCart, cartID uuid.UUID) (dto.UpdateCart, error)
	GetCartByID(cartID uuid.UUID) (dto.GetCartByCartID, error)
	DeleteCartByCartID(CartID uuid.UUID) (dto.DeleteCartByCartID, error)
	DeleteCartByUserID(UserID uuid.UUID) (dto.DeleteCartByUserID, error)
}

type CartUsecase struct {
	cartRepo repository.CartMySQLItf
}

func NewCartUsecase(cartRepo repository.CartMySQLItf) CartUsecaseItf {
	return &CartUsecase{
		cartRepo: cartRepo,
	}
}

func (u *CartUsecase) CreateCart(cart dto.CreateCart, userID uuid.UUID, sellerID uuid.UUID) (dto.CreateCart, error) {
	cartUser := entity.Cart{
		CartItemID: uuid.New(),
		UserID:     userID,
		ProductID:  cart.ProductID,
		SellerID:   sellerID,
		Count:      cart.Count,
	}

	err := u.cartRepo.CreateCart(&cartUser)
	if err != nil {
		return dto.CreateCart{}, fiber.NewError(http.StatusInternalServerError, "failed to create cart")
	}

	return cartUser.ParseToDTOCreateCart(), nil
}

func (u *CartUsecase) UpdateCart(cart dto.UpdateCart, cartID uuid.UUID) (dto.UpdateCart, error) {
	cartUser := entity.Cart{
		CartItemID: cartID,
		Count:      cart.Count,
	}

	err := u.cartRepo.UpdateCart(&cartUser)
	if err != nil {
		return dto.UpdateCart{}, fiber.NewError(http.StatusInternalServerError, "failed to update cart")
	}

	return cartUser.ParseToDTOUpdateCart(), nil
}

func (u *CartUsecase) GetCartByID(cartID uuid.UUID) (dto.GetCartByCartID, error) {
	cartUser := entity.Cart{
		CartItemID: cartID,
	}

	err := u.cartRepo.GetCartByID(&cartUser)
	if err != nil {
		return dto.GetCartByCartID{}, fiber.NewError(http.StatusInternalServerError, "failed to get cart by id")
	}

	return cartUser.ParseToDTOGetCartByCartID(), nil
}

func (u *CartUsecase) DeleteCartByCartID(CartID uuid.UUID) (dto.DeleteCartByCartID, error) {
	cartUser := entity.Cart{
		CartItemID: CartID,
	}

	err := u.cartRepo.DeleteCartByCartID(&cartUser)
	if err != nil {
		return dto.DeleteCartByCartID{}, fiber.NewError(http.StatusInternalServerError, "failed to delete cart by id")
	}

	return cartUser.ParseToDTODeleteCartByCartID(), nil
}

func (u *CartUsecase) DeleteCartByUserID(UserID uuid.UUID) (dto.DeleteCartByUserID, error) {
	cartUserID := entity.Cart{
		UserID: UserID,
	}

	err := u.cartRepo.DeleteCartByUserID(&cartUserID, UserID)
	if err != nil {
		return dto.DeleteCartByUserID{}, fiber.NewError(http.StatusInternalServerError, "failed to delete cart by user id")
	}

	return cartUserID.ParseToDTODeleteCartByUserID(), nil
}
