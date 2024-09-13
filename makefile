
hello:
	echo "Hello"

build:
	go build -o bin/main main.go

run:
	go run main.go
	

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go


all: hello build
