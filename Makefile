run:
	go run main.go
build:
	go build main.go
clean:
	go mod tidy
swag:
	swag init -g main.go --output docs