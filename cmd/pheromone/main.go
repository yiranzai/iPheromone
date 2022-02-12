package pheromone

import (
	"context"

	"github.com/yiranzai/iPheromone/app/application"
	"github.com/yiranzai/iPheromone/app/interfaces"
)

func main() {

	// 依赖注入
	s := interfaces.NewSearchInf(&application.ImgSearchApp{})

	err := s.Search(context.Background(), "keywords", "")
	if err != nil {
		return
	}
}
