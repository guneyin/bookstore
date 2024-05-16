BINARY_NAME=bookstore

build:
	go build -o ${BINARY_NAME} cmd/app/main.go

run:
	go run cmd/app/main.go

clean:
	go clean
	rm ${BINARY_NAME}