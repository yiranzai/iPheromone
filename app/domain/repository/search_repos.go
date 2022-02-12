package repository

import (
	"github.com/yiranzai/golang-project-template/app/domain/entity"
)

type SearchImage struct {
}

type IReposSearch interface {
	Search(req *entity.SearchKeywords) []*SearchImage
	// SearchTags(req *entity.SearchTagReq) []*SearchImage
}
