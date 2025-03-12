package usecase

import (
	"lentara-backend/internal/app/cart/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CartUseCaseItf interface {
	CreateCart(cart dto.CreateCart, userID uuid.UUID, sellerID uuid.UUID) (dto.CreateCart, error)
	UpdateCart(cart dto.UpdateCart, cartID uuid.UUID) (dto.UpdateCart, error)
	GetCartByID(cartID uuid.UUID) (dto.GetCartByCartID, error)
	GetCartsByUserID(user uuid.UUID) (*[]dto.GetCartsByUserID, error)
	DeleteCartByCartID(CartID uuid.UUID) (dto.DeleteCartByCartID, error)
	DeleteCartByUserID(UserID uuid.UUID) (dto.DeleteCartByUserID, error)
	GetSellerListFromUserCart(userID uuid.UUID) ([]string, error)
}

type CartUseCase struct {
	cartRepo repository.CartMySQLItf
}

func NewCartUseCase(cartRepo repository.CartMySQLItf) CartUseCaseItf {
	return &CartUseCase{
		cartRepo: cartRepo,
	}
}

func (u CartUseCase) CreateCart(cart dto.CreateCart, userID uuid.UUID, sellerID uuid.UUID) (dto.CreateCart, error) {
	cartUser := entity.Cart{
		CartItemID:   uuid.New(),
		UserID:       userID,
		ProductID:    cart.ProductID,
		SellerID:     sellerID,
		Count:        cart.Count,
		RentDuration: cart.RentDuration,
	}

	err := u.cartRepo.CreateCart(&cartUser)
	if err != nil {
		return dto.CreateCart{}, fiber.NewError(http.StatusInternalServerError, "failed to create cart")
	}

	return cartUser.ParseToDTOCreateCart(), nil
}

func (u CartUseCase) UpdateCart(cart dto.UpdateCart, cartID uuid.UUID) (dto.UpdateCart, error) {
	cartUser := entity.Cart{
		CartItemID:   cartID,
		Count:        cart.Count,
		RentDuration: cart.RentDuration,
	}

	err := u.cartRepo.UpdateCart(&cartUser)
	if err != nil {
		return dto.UpdateCart{}, fiber.NewError(http.StatusInternalServerError, "failed to update cart")
	}

	return cartUser.ParseToDTOUpdateCart(), nil
}

func (u CartUseCase) GetCartByID(cartID uuid.UUID) (dto.GetCartByCartID, error) {
	cartUser := entity.Cart{
		CartItemID: cartID,
	}

	err := u.cartRepo.GetCartByID(&cartUser)
	if err != nil {
		return dto.GetCartByCartID{}, fiber.NewError(http.StatusInternalServerError, "failed to get cart by id")
	}

	return cartUser.ParseToDTOGetCartByCartID(), nil
}

func (u CartUseCase) GetCartsByUserID(user uuid.UUID) (*[]dto.GetCartsByUserID, error) {
	cartUserResult := new([]entity.Cart)

	err := u.cartRepo.GetCartsByUserID(cartUserResult, user)
	if err != nil {
		return nil, fiber.NewError(http.StatusInternalServerError, "failed to get carts by user id")
	}

	res := make([]dto.GetCartsByUserID, len(*cartUserResult))
	for i, cart := range *cartUserResult {
		res[i] = cart.ParseToDTOGetCartsByUserID()
	}

	return &res, nil
}

func (u CartUseCase) GetCartsByUserAndSellerID(userID uuid.UUID, sellerID uuid.UUID) (*[]dto.GetCartsByUserIDAndSellerID, error) {
	cartUserResult := new([]entity.Cart)

	err := u.cartRepo.GetCartsByUserIDAndSellerID(cartUserResult, userID, sellerID)
	if err != nil {
		return nil, fiber.NewError(http.StatusInternalServerError, "failed to get carts by user id and seller id")
	}

	res := make([]dto.GetCartsByUserIDAndSellerID, len(*cartUserResult))
	for i, cart := range *cartUserResult {
		res[i] = cart.ParseToDTOGetCartsByUserIDAndSellerID()
	}

	return &res, nil
}

func (u CartUseCase) DeleteCartByCartID(CartID uuid.UUID) (dto.DeleteCartByCartID, error) {
	cartUser := entity.Cart{
		CartItemID: CartID,
	}

	err := u.cartRepo.DeleteCartByCartID(&cartUser)
	if err != nil {
		return dto.DeleteCartByCartID{}, fiber.NewError(http.StatusInternalServerError, "failed to delete cart by id")
	}

	return cartUser.ParseToDTODeleteCartByCartID(), nil
}

func (u CartUseCase) DeleteCartByUserID(UserID uuid.UUID) (dto.DeleteCartByUserID, error) {
	cartUserID := entity.Cart{
		UserID: UserID,
	}

	err := u.cartRepo.DeleteCartByUserID(&cartUserID, UserID)
	if err != nil {
		return dto.DeleteCartByUserID{}, fiber.NewError(http.StatusInternalServerError, "failed to delete cart by user id")
	}

	return cartUserID.ParseToDTODeleteCartByUserID(), nil
}

func (u CartUseCase) GetSellerListFromUserCart(userID uuid.UUID) ([]string, error) {
	cartUserResult := new([]string)

	err := u.cartRepo.GetSellerListFromUserCart(cartUserResult, userID)
	if err != nil {
		return nil, fiber.NewError(http.StatusInternalServerError, "failed to get seller list from user id")
	}

	return *cartUserResult, nil
}
