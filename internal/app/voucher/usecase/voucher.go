package usecase

import (
	"lentara-backend/internal/app/voucher/repository"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type VoucherUseCaseItf interface {
	CreateVoucher(voucher dto.CreateVoucher) (dto.ResponseCreateVoucher, error)
	UpdateVoucher(voucher dto.UpdateVoucher, voucherID uuid.UUID) (dto.UpdateVoucher, error)
	GetVoucherByID(voucherID uuid.UUID) (dto.GetVoucherByID, error)
	GetAllVouchers() (*[]dto.GetAllVouchers, error)
}

type VoucherUseCase struct {
	voucherRepo repository.VoucherMySQLItf
}

func NewVoucherUseCase(voucherRepo repository.VoucherMySQLItf) VoucherUseCaseItf {
	return &VoucherUseCase{
		voucherRepo: voucherRepo,
	}
}

func (u VoucherUseCase) CreateVoucher(voucher dto.CreateVoucher) (dto.ResponseCreateVoucher, error) {
	createVoucher := entity.Voucher{
		ID:    uuid.New(),
		Title: voucher.Title,
		Count: voucher.Count,
		Value: voucher.Value,
		Type:  voucher.Type,
	}

	err := u.voucherRepo.CreateVoucher(&createVoucher)
	if err != nil {
		return dto.ResponseCreateVoucher{}, fiber.NewError(http.StatusInternalServerError, "failed to create new voucher")
	}

	return createVoucher.ParseToDTOResponseCreateVoucher(), nil
}

func (u VoucherUseCase) UpdateVoucher(voucher dto.UpdateVoucher, voucherID uuid.UUID) (dto.UpdateVoucher, error) {
	updateVoucher := entity.Voucher{
		ID:    voucherID,
		Title: voucher.Title,
		Count: voucher.Count,
		Value: voucher.Value,
		Type:  voucher.Type,
	}

	err := u.voucherRepo.UpdateVoucher(&updateVoucher)
	if err != nil {
		return dto.UpdateVoucher{}, fiber.NewError(http.StatusInternalServerError, "failed to update voucher")
	}

	return updateVoucher.ParseToDTOUpdateVoucher(), nil
}

func (u VoucherUseCase) GetVoucherByID(voucherID uuid.UUID) (dto.GetVoucherByID, error) {
	voucher := entity.Voucher{
		ID: voucherID,
	}

	err := u.voucherRepo.GetVoucherByID(&voucher, voucherID)
	if err != nil {
		return dto.GetVoucherByID{}, fiber.NewError(http.StatusInternalServerError, "failed to get voucher by id")
	}

	return voucher.ParseToDTOGetVoucherByID(), nil
}

func (u VoucherUseCase) GetAllVouchers() (*[]dto.GetAllVouchers, error) {
	voucher := new([]entity.Voucher)

	err := u.voucherRepo.GetAllVouchers(voucher)
	if err != nil {
		return nil, fiber.NewError(http.StatusInternalServerError, "failed to get all vouchers")
	}

	res := make([]dto.GetAllVouchers, len(*voucher))
	for i, cur := range *voucher {
		res[i] = cur.ParseToDTOGetAllVouchers()
	}

	return &res, nil
}
