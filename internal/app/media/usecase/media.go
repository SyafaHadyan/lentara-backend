package usecase

import "lentara-backend/internal/app/media/repository"

type MediaUseCaseItf interface{}

type MediaUseCase struct {
	MediaRepository repository.MediaMySQLItf
}

func NewMediaUseCase(mediaRepository repository.MediaMySQLItf) MediaUseCaseItf {
	return &MediaUseCase{
		MediaRepository: mediaRepository,
	}
}
