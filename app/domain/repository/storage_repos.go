package repository

import (
	"github.com/yiranzai/golang-project-template/app/domain/entity"
)

type IReposStorage interface {
	SaveImage(imgs []*entity.UploadImage) ([]*entity.UploadImage, error)
	FindImage(ids []uint64) ([]*entity.UploadImage, error)
}
