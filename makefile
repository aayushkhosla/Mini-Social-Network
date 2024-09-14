
hello:
	echo "Hello"

build:
	go build -o bin/main main.go

run:
	CompileDaemon -directory=./  -command="go run main.go"CompileDaemon -directory=./  -command="go run main.go"
compile:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go
reset:
	./goose down
	./goose down
	./goose down
	./goose down
	./goose up 


all: hello build
