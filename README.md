# go-web

Gin demo

## tech stack

* package management using `go module`
* http framework using `gin`
* mysql driver using `gorm`
* mongodb driver using mongo
* load env using `godotenv`
* load config file using `viper`

## basic

### package manage

using go module, `go version >= 1.11` , reference: https://blog.golang.org/using-go-modules

``` 
// install package
go get github.com/spf13/viper

// remove all packages that not using
go mod tidy
```
