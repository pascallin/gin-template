FROM golang:1.18

WORKDIR /go/src/app
COPY . .

# NOTE: skip env config for now
RUN go mod download
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init
RUN go build -o ./bin/gin-server
RUN go install -v .

CMD ["gin-server"]