package usecase

import (
	"lentara-backend/internal/app/payment/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PaymentUseCaseItf interface {
	StorePayment(payment dto.StorePayment, orderID uuid.UUID, userID uuid.UUID, totalPrice uint64) (dto.StorePayment, error)
	UpdatePayment(payment dto.MidtransUpdatePaymentStatus) (dto.UpdatePaymentStatus, error)
	GetPaymentInfo(userID uuid.UUID, orderID uuid.UUID) (dto.GetPaymentStatus, error)
	GetUserIDFromOrderID(orderID uuid.UUID) (uuid.UUID, error)
}

type PaymentUseCase struct {
	paymentRepo repository.PaymentMySQLItf
}

func NewPaymentUseCase(paymentRepo repository.PaymentMySQLItf) PaymentUseCaseItf {
	return &PaymentUseCase{
		paymentRepo: paymentRepo,
	}
}

func (u PaymentUseCase) StorePayment(payment dto.StorePayment, orderID uuid.UUID, userID uuid.UUID, totalPrice uint64) (dto.StorePayment, error) {
	paymentUser := entity.Payment{
		ID:         orderID,
		UserID:     userID,
		TotalPrice: totalPrice,
		Status:     "pending",
	}

	err := u.paymentRepo.StorePayment(&paymentUser)
	if err != nil {
		return dto.StorePayment{}, fiber.NewError(http.StatusInternalServerError, "failed to store new payment")
	}

	return paymentUser.ParseToDTOStorePayment(), nil
}

func (u PaymentUseCase) UpdatePayment(payment dto.MidtransUpdatePaymentStatus) (dto.UpdatePaymentStatus, error) {
	orderID, err := uuid.Parse(payment.OrderID)
	if err != nil {
		return dto.UpdatePaymentStatus{}, fiber.NewError(http.StatusBadRequest, "invalid order id")
	}

	paymentUser := entity.Payment{
		ID:     orderID,
		Status: payment.TransactionStatus,
	}

	err = u.paymentRepo.UpdatePayment(&paymentUser)
	if err != nil {
		return dto.UpdatePaymentStatus{}, fiber.NewError(http.StatusInternalServerError, "failed to update payment status")
	}

	return paymentUser.ParseToDTOUpdatePayment(), nil
}

func (u PaymentUseCase) GetPaymentInfo(userID uuid.UUID, orderID uuid.UUID) (dto.GetPaymentStatus, error) {
	orderUser := entity.Payment{
		ID:     orderID,
		UserID: userID,
	}

	err := u.paymentRepo.GetPaymentInfo(&orderUser, userID, orderID)
	if err != nil {
		return dto.GetPaymentStatus{}, fiber.NewError(http.StatusInternalServerError, "failed to get payment info")
	}

	return orderUser.ParseToDTOGetPaymentStatus(), nil
}

func (u PaymentUseCase) GetUserIDFromOrderID(orderID uuid.UUID) (uuid.UUID, error) {
	orderInfo := entity.Payment{
		ID: orderID,
	}

	err := u.paymentRepo.GetPaymentUserInfo(&orderInfo, orderID)
	if err != nil {
		return uuid.Nil, fiber.NewError(http.StatusInternalServerError, "failed to get user info")
	}

	return orderInfo.UserID, nil
}
