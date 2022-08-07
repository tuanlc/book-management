BINARY_NAME=main.out

deps:
	go mod download

test:
	go test ./... -v

dev:
	air

build:
	go build -o ${BINARY_NAME} ./cmd/main.go

run:
	go build -o ${BINARY_NAME} ./cmd/main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}