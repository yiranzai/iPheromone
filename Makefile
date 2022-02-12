PROJECT="iPheromone"

default:
	go build -o ./bin

.PHONY:run
run:
	go run main.go

.PHONY:linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin

.PHONY:osx
osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin

.PHONY:windows
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin

.PHONY:help
help:
	@echo ""
	@echo "Usage:"
	@echo "\t make [platform]"
	@echo "\t 默认编译当前平台版本"
	@echo "支持平台列表"
	@echo "\t windows \t Windows平台"
	@echo "\t linux \t\t Linux平台"
	@echo "\t osx \t\t MacOS平台"
	@echo "\t run \t\t 启动项目"
