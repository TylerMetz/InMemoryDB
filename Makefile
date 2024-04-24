SRC_DIR := src
BUILD_DIR := bin

build:
	go build -o $(BUILD_DIR)/InMemoryDB.out ./$(SRC_DIR)

run: build
	go run $(SRC_DIR)/main.go
	