
deps:
	@echo "### Pulling dependency libraries."
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/lib/pq
	go get -u github.com/gin-gonic/gin
	go get -u github.com/go-redis/redis
	go get -u github.com/swaggo/swag/cmd/swag
	go get -u github.com/gin-contrib/location
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/gin-swagger/swaggerFiles
	go get -u github.com/alecthomas/template

build:
	@echo "### Building taco-registry Builder."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --ldflags=--s -o builder src/builder/main.go

docker-build:
	@echo "### Making Builder docker image."
	docker build --network=host --no-cache -t taco-registry/builder:latest . -f ./Dockerfile.multi

docker-build-single:
	@echo "### Making Builder docker image."
	docker build --network=host --no-cache -t taco-registry/builder:latest . -f ./Dockerfile