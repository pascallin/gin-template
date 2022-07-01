FROM golang:1.18

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.3
RUN swag init
RUN go build -o ./bin/gin-server
RUN go install -v .

CMD ["gin-server"]