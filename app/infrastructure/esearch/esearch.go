package esearch

import (
	"github.com/yiranzai/iPheromone/app/domain/entity"
	"github.com/yiranzai/iPheromone/app/domain/repository"
)

type ESearch struct {
}

func (e *ESearch) Search(req *entity.SearchKeywords) []*repository.SearchImage {
	panic("implement me")
}
