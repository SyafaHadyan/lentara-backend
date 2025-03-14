package rest

import (
	"fmt"
	usecase "lentara-backend/internal/app/media/usecase"
	productusecase "lentara-backend/internal/app/product/usecase"
	"lentara-backend/internal/infra/env"
	"lentara-backend/internal/middleware"
	"net/http"

	// supabaseupload "github.com/adityarizkyramadhan/supabase-storage-uploader"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	// supabasestorage "github.com/supabase-community/storage-go"
)

type MediaHandler struct {
	Validator      *validator.Validate
	Config         *env.Env
	Middleware     middleware.MiddlewareItf
	MediaUseCase   usecase.MediaUseCaseItf
	ProductUseCase productusecase.ProductUseCaseItf
}

func NewMediaHandler(routerGroup fiber.Router, validator *validator.Validate, config *env.Env, middleware middleware.MiddlewareItf, mediaUseCase usecase.MediaUseCaseItf, productUseCase productusecase.ProductUseCaseItf) {
	handler := MediaHandler{
		Validator:      validator,
		Config:         config,
		Middleware:     middleware,
		MediaUseCase:   mediaUseCase,
		ProductUseCase: productUseCase,
	}

	routerGroup = routerGroup.Group("/media")

	routerGroup.Post("/upload", handler.UploadMedia)
}

func (h MediaHandler) UploadMedia(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "failed to get media")
	}

	fileName := uuid.New().String() + ".jpg"

	ctx.SaveFile(file, fmt.Sprintf("../%s", fileName))

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "successfully uploaded media",
		"payload": "https://lentara-be.syafahadyan.com/images/" + fileName,
	})

	// os.WriteFile("file", file)

	// bucketName := h.Config.AWSS3Bucket
	// s3Region := h.Config.AWSS3Region
	// uploadDir := "/"
	// var s3Client *s3.Client

	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(s3Region))
	// if err != nil {
	// 	return fiber.NewError(http.StatusInternalServerError, "failed to fetch bucket info")
	// }

	// s3Client = s3.NewFromConfig(cfg)

	// s3Key := filepath.Join(uploadDir, )

	// _, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: aws.String(bucketName),
	// 	Key:    aws.String(s3Key),
	// 	Body:   ,
	// })
	// if err != nil {
	// 	return fiber.NewError(http.StatusInternalServerError, "failed to upload file to bucket")
	// }

	// storageClient := supabasestorage.NewClient(h.Config.SupabaseProjectURL, h.Config.SupabaseAccessToken, nil)

	// supClient := supabaseupload.New(
	// 	h.Config.SupabaseProjectURL,
	// 	h.Config.SupabaseAPIKey,
	// 	h.Config.SupabaseBukcetName,
	// )

	// file, err := ctx.FormFile("file")
	// if err != nil {
	// 	return fiber.NewError(http.StatusBadRequest, "failed to get media")
	// }

	// res, err := storageClient.GetBucket(h.Config.SupabaseBukcetName)
	// if err != nil {
	// 	log.Println(err)
	// }

	// result, err := storageClient.UploadFile("test", "test.txt", file)

	// link, err := supClient.Upload(file)
	// if err != nil {
	// 	return fiber.NewError(http.StatusInternalServerError, "failed to upload media")
	// }

	// return ctx.Status(http.StatusCreated).JSON(fiber.Map{
	// "message": "successfully uploaded media",
	// "payload": res,
	// "upload":  result,
	// })
}
