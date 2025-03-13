package usecase

import (
	"lentara-backend/internal/app/cart/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"lentara-backend/internal/infra/env"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CartUseCaseItf interface {
	CreateCart(cart dto.CreateCart, userID uuid.UUID, sellerID uuid.UUID, price uint64) (dto.CreateCart, error)
	UpdateCart(cart dto.UpdateCart, cartID uuid.UUID) (dto.UpdateCart, error)
	GetCartByID(cartID uuid.UUID) (dto.GetCartByCartID, error)
	GetCartsByUserID(user uuid.UUID) (*[]dto.GetCartsByUserID, error)
	DeleteCartByCartID(CartID uuid.UUID) (dto.DeleteCartByCartID, error)
	DeleteCartByUserID(UserID uuid.UUID) (dto.DeleteCartByUserID, error)
	GetSellerListFromUserCart(userID uuid.UUID) ([]string, error)
	GetCartSummary(userID uuid.UUID) (dto.GetCartSummary, error)
}

type CartUseCase struct {
	cartRepo repository.CartMySQLItf
	config   *env.Env
}

func NewCartUseCase(cartRepo repository.CartMySQLItf, config *env.Env) CartUseCaseItf {
	return &CartUseCase{
		cartRepo: cartRepo,
		config:   config,
	}
}

func (u CartUseCase) CreateCart(cart dto.CreateCart, userID uuid.UUID, sellerID uuid.UUID, price uint64) (dto.CreateCart, error) {
	cartUser := entity.Cart{
		CartItemID:   uuid.New(),
		UserID:       userID,
		ProductID:    cart.ProductID,
		SellerID:     sellerID,
		Count:        cart.Count,
		Price:        price,
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

func (u CartUseCase) GetCartSummary(userID uuid.UUID) (dto.GetCartSummary, error) {
	cartUserResult := new([]entity.Cart)

	err := u.cartRepo.GetCartsByUserID(cartUserResult, userID)
	if err != nil {
		return dto.GetCartSummary{}, fiber.NewError(http.StatusInternalServerError, "failed to get carts from user id")
	}

	var productCount uint8
	var totalPrice uint64
	var voucher uint64
	var serviceCost uint64
	var DepositeAmount uint64

	for _, cart := range *cartUserResult {
		productCount += cart.Count
		totalPrice += cart.Price
	}

	serviceCost = uint64(float64(totalPrice) * float64(u.config.SerivceCost) / float64(100))
	DepositeAmount = uint64(float64(totalPrice) * float64(u.config.DepositePercentage) / float64(100))
	totalPrice += (serviceCost + DepositeAmount - voucher)

	cartSummary := dto.GetCartSummary{
		UserID:             userID,
		ProductCount:       productCount,
		DeliveryCost:       0,
		ServiceCost:        serviceCost,
		DepositeAmount:     DepositeAmount,
		DepositePercentage: uint64(u.config.DepositePercentage),
		Voucher:            0,
		TotalPrice:         totalPrice,
	}

	return cartSummary, nil
}
