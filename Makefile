BIN_DIR=bin
SRC_DIR=cmd


.PHONY: build
build:
	go build -o $(BIN_DIR)/$(c) $(SRC_DIR)/$(c)/main.go

.PHONY: run
run: build
	./$(BIN_DIR)/$(c)