# go-web

Gin server Template for quick start

## tech stack

* package management using `go module`
* http framework using [gin](https://github.com/gin-gonic/gin)
* mysql driver using [gorm](https://github.com/go-gorm/gorm)
* mongodb driver using [mongo](https://github.com/mongodb/mongo-go-driver)
* load env using [godotenv](https://github.com/joho/godotenv)
* using [gin-swagger](https://github.com/swaggo/gin-swagger) to generate swagger docs
* hot-reload tool using [air](https://github.com/cosmtrek/air)

## package manage

using go module, `go version >= 1.18`, reference: <https://blog.golang.org/using-go-modules>

```shell
go mod download
```

## swagger

1. generate swagger json file in project root folder

```shell
go install github.com/swaggo/swag/cmd/swag@latest
swag init --generalInfo ./server/server.go
```

1. visit <http://localhost:4000/swagger/index.html>

reference: <https://github.com/swaggo/swag/blob/master/README_zh-CN.md>

## run as docker container

```shell
docker build -t my-gin-server .
docker run -it --rm --name my-running-gin-server my-gin-server
```

## build and run

```shell
go build -o ./bin/gin-server
go install -v .
gin-server
```

## testing

```shell
# install command
go get github.com/joho/godotenv/cmd/godotenv
# running test with specific env
godotenv -f .env.test go test ./...
```
