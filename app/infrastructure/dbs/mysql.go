package dbs

import (
	"github.com/yiranzai/golang-project-template/app/domain/entity"
)

type Dbs struct {
}

func (d *Dbs) SaveImage(imgs []*entity.UploadImage) ([]*entity.UploadImage, error) {
	panic("implement me")
}

func (d *Dbs) FindImage(ids []uint64) ([]*entity.UploadImage, error) {
	panic("implement me")
}
