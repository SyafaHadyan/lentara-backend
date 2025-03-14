package rest

import (
	cartusecase "lentara-backend/internal/app/cart/usecase"
	usecase "lentara-backend/internal/app/payment/usecase"
	productusecase "lentara-backend/internal/app/product/usecase"
	sellerusecase "lentara-backend/internal/app/seller/usecase"
	userusecase "lentara-backend/internal/app/user/usecase"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/infra/env"
	"lentara-backend/internal/middleware"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentHandler struct {
	Validator      *validator.Validate
	Config         *env.Env
	Middleware     middleware.MiddlewareItf
	PaymentUseCase usecase.PaymentUseCaseItf
	ProductUseCase productusecase.ProductUseCaseItf
	UserUseCase    userusecase.UserUseCaseItf
	SellerUseCase  sellerusecase.SellerUseCaseItf
	CartUseCase    cartusecase.CartUseCaseItf
	Midtrans       midtrans.HttpClient
}

func NewPaymentHandler(routerGroup fiber.Router, validator *validator.Validate, config *env.Env, middleware middleware.MiddlewareItf, paymentUseCase usecase.PaymentUseCaseItf, productUseCase productusecase.ProductUseCaseItf, userUseCase userusecase.UserUseCaseItf, sellerUseCase sellerusecase.SellerUseCaseItf, cartUseCase cartusecase.CartUseCaseItf) {
	paymentHandler := PaymentHandler{
		Validator:      validator,
		Config:         config,
		Middleware:     middleware,
		PaymentUseCase: paymentUseCase,
		ProductUseCase: productUseCase,
		SellerUseCase:  sellerUseCase,
		UserUseCase:    userUseCase,
		CartUseCase:    cartUseCase,
	}

	routerGroup = routerGroup.Group("/payment")

	routerGroup.Post("/new", middleware.Authentication, paymentHandler.StorePayment)
	routerGroup.Post("/update", paymentHandler.UpdatePayment)
	routerGroup.Get("/status/:id", middleware.Authentication, paymentHandler.GetPaymentStatus)
}

func GenerateSnapReq(orderID uuid.UUID, grossAmt int64) *snap.Request {
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID.String(),
			GrossAmt: grossAmt,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	return snapReq
}

func (h PaymentHandler) StorePayment(ctx *fiber.Ctx) error {
	var store dto.StorePayment
	var s snap.Client
	orderID := uuid.New()
	midtrans.ServerKey = h.Config.MidtransServerKey
	var midtransEnv midtrans.EnvironmentType
	switch h.Config.MidtransEnvironment {
	case 0:
		midtransEnv = midtrans.Sandbox
	case 1:
		midtransEnv = midtrans.Production
	}

	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unathorized")
	}

	userCart, err := h.CartUseCase.GetOrderSummary(userID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get user cart summary")
	}

	res, err := h.PaymentUseCase.StorePayment(store, orderID, userID, userCart.TotalPrice)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to store new payment")
	}

	log.Println(midtrans.Environment)

	s.New(midtrans.ServerKey, midtransEnv)

	snapReq := GenerateSnapReq(orderID, int64(userCart.TotalPrice))

	mtRes, _ := s.CreateTransaction(snapReq)

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message":  "successfully stored new payment",
		"payload":  res,
		"midtrans": mtRes,
	})
}

func (h PaymentHandler) UpdatePayment(ctx *fiber.Ctx) error {
	var update dto.MidtransUpdatePaymentStatus

	err := ctx.BodyParser(&update)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	res, err := h.PaymentUseCase.UpdatePayment(update)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update payment status")
	}

	orderID, err := uuid.Parse(update.OrderID)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid user id")
	}

	userID, err := h.PaymentUseCase.GetUserIDFromOrderID(orderID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get user info")
	}

	var cartDeletetionStatus dto.DeleteCartByUserID

	if res.Status == "capture" || res.Status == "settlement" {
		cartDeletetionStatus, err = h.CartUseCase.DeleteCartByUserID(userID)
		if err != nil {
			log.Println("failed to delete cart from user id")
		}
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":               "successfully updated payment status",
		"cart_deleteion_status": cartDeletetionStatus,
		"payload":               res,
	})
}

func (h PaymentHandler) GetPaymentStatus(ctx *fiber.Ctx) error {
	userID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unathorized")
	}

	orderID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid order id")
	}

	res, err := h.PaymentUseCase.GetPaymentInfo(userID, orderID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get payment info")
	}

	var cartDeletetionStatus dto.DeleteCartByUserID

	if res.Status == "capture" || res.Status == "settlement" {
		cartDeletetionStatus, err = h.CartUseCase.DeleteCartByUserID(userID)
		if err != nil {
			log.Println("failed to delete cart from user id")
		}
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":               "successfully get payment status",
		"cart_deleteion_status": cartDeletetionStatus,
		"payload":               res,
	})
}
