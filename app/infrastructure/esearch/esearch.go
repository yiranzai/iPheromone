package esearch

import (
	"github.com/yiranzai/golang-project-template/app/domain/entity"
	"github.com/yiranzai/golang-project-template/app/domain/repository"
)

type ESearch struct {
}

func (e *ESearch) Search(req *entity.SearchKeywords) []*repository.SearchImage {
	panic("implement me")
}
