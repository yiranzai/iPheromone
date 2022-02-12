package myapp

import (
	"context"

	"github.com/yiranzai/golang-project-template/app/application"
	"github.com/yiranzai/golang-project-template/app/interfaces"
)

func main() {

	// 依赖注入
	s := interfaces.NewSearchInf(&application.ImgSearchApp{})

	err := s.Search(context.Background(), "keywords", "")
	if err != nil {
		return
	}
}
