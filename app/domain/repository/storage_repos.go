package repository

import (
	"github.com/yiranzai/iPheromone/app/domain/entity"
)

type IReposStorage interface {
	SaveImage(imgs []*entity.UploadImage) ([]*entity.UploadImage, error)
	FindImage(ids []uint64) ([]*entity.UploadImage, error)
}
