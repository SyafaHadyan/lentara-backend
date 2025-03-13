package rest

import (
	usecase "lentara-backend/internal/app/product/usecase"
	sellerusecase "lentara-backend/internal/app/seller/usecase"
	"lentara-backend/internal/domain/dto"
	"lentara-backend/internal/domain/entity"
	"lentara-backend/internal/middleware"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct {
	Validator      *validator.Validate
	Middleware     middleware.MiddlewareItf
	ProductUseCase usecase.ProductUseCaseItf
	SellerUseCase  sellerusecase.SellerUseCaseItf
}

func NewProductHandler(routerGroup fiber.Router, validator *validator.Validate, middleware middleware.MiddlewareItf, productUseCase usecase.ProductUseCaseItf, sellerUseCase sellerusecase.SellerUseCaseItf) {
	handler := ProductHandler{
		Validator:      validator,
		Middleware:     middleware,
		ProductUseCase: productUseCase,
		SellerUseCase:  sellerUseCase,
	}

	routerGroup = routerGroup.Group("/")

	routerGroup.Get("/products/", handler.GetAllProducts)
	routerGroup.Get("/products/:id", handler.GetProductByID)
	routerGroup.Get("/seller/products/:id", handler.GetProductsBySellerID)
	routerGroup.Get("/products/category/:category", handler.GetProductCategory)
	routerGroup.Get("/products/sortprice/:price", handler.SortProductByPrice)
	routerGroup.Get("/products/rating/:rating", handler.FilterProductByRating)
	routerGroup.Get("/search/:title", handler.SearchProduct)
	routerGroup.Post("/products", middleware.Authentication, handler.CreateProduct)
	routerGroup.Patch("/products/:id", middleware.Authentication, handler.UpdateProduct)
	routerGroup.Delete("/products/:id", middleware.Authentication, middleware.AdminUser, handler.DeleteProduct)
}

func (h ProductHandler) GetAllProducts(ctx *fiber.Ctx) error {
	//page, err := strconv.Atoi(ctx.Params("page"))
	//if err != nil {
	//	return fiber.NewError(http.StatusBadRequest, "page number must be specified")
	//}

	res, err := h.ProductUseCase.GetAllProducts()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get products")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}

func (h ProductHandler) GetProductByID(ctx *fiber.Ctx) error {
	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "not a valid uuid")
	}

	product, err := h.ProductUseCase.GetProductByID(productID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get product")
	}

	seller, err := h.SellerUseCase.GetSellerInfo(product.SellerID)
	if err != nil {
		// Continue even if failed to get seller info
		log.Println("failed to get seller info")
		// return fiber.NewError(http.StatusInternalServerError, "failed to get seller info")
	}

	sellerInfo := &entity.Seller{
		ID:             seller.ID,
		Name:           seller.Name,
		StoreName:      seller.StoreName,
		StoreLocation:  seller.StoreLocation,
		ProfilePicture: seller.ProfilePicture,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "successfully get product by id",
		"product": product,
		"seller":  sellerInfo.ParseToDTOGetPublicSellerInfo(),
	})
}

func (h ProductHandler) GetProductsBySellerID(ctx *fiber.Ctx) error {
	sellerID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid seller id")
	}

	res, err := h.ProductUseCase.GetProductsBySellerID(sellerID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get products by seller id")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}

func (h ProductHandler) GetProductCategory(ctx *fiber.Ctx) error {
	res, err := h.ProductUseCase.GetProductCategory(ctx.Params("category"))
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get product category")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}

func (h ProductHandler) SortProductByPrice(ctx *fiber.Ctx) error {
	sort := ctx.Params("price")

	res, err := h.ProductUseCase.SortProductByPrice(sort)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to sort product by price")
	}

	var message string
	if sort != "lowest" && sort != "highest" {
		message = "invalid sort request, returning unsorted products"
	} else {
		message = "successfully sorted products from " + sort + " price"
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": message,
		"payload": res,
	})
}

func (h ProductHandler) FilterProductByRating(ctx *fiber.Ctx) error {
	rating := ctx.Params("rating")

	res, err := h.ProductUseCase.FilterProductByRating(rating)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to sort product by price")
	}

	var message string
	if len(rating) == 0 {
		message = "invalid rating param, returning products using default filter"
	} else {
		message = "successfully sorted products with rating " + rating
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": message,
		"payload": res,
	})
}

func (h ProductHandler) SearchProduct(ctx *fiber.Ctx) error {
	res, err := h.ProductUseCase.SearchProduct(ctx.Params("title"))
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to search product")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}

func (h ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	var request dto.RequestCreateProduct

	err := ctx.BodyParser(&request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to parse request body")
	}

	log.Println(request)

	err = h.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to validate request")
	}

	sellerID, err := uuid.Parse(ctx.Locals("userID").(string))
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "user unathorized")
	}

	if sellerID == uuid.Nil {
		return fiber.NewError(http.StatusUnauthorized, "user unathorized")
	}

	productOrigin, err := h.SellerUseCase.GetSellerInfo(sellerID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get seller info")
	}

	res, err := h.ProductUseCase.CreateProduct(request, sellerID, productOrigin.StoreLocation)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to create product")
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "successfully created prodcut",
		"payload": res,
	})
}

func (h ProductHandler) UpdateProduct(ctx *fiber.Ctx) error {
	var request dto.UpdateProduct

	err := ctx.BodyParser(&request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update product")
	}

	productID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "can't find product with current id")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "payload invalid")
	}

	err = h.ProductUseCase.UpdateProduct(productID, request)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to update product")
	}

	product, err := h.ProductUseCase.GetProductByID(productID)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get product info")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "product udpated",
		"payload": product,
	})
}

func (h ProductHandler) DeleteProduct(ctx *fiber.Ctx) error {
	var request dto.DeleteProduct

	err := ctx.BodyParser(request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to parse request with current id")
	}

	ProductID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get product id")
	}

	err = h.Validator.Struct(request)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid product id")
	}

	res, err := h.ProductUseCase.DeleteProduct(ProductID, request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to delete product")
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"payload": res,
	})
}
