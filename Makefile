BIN_DIR:=bin

.PHONY: build clean test

build:
	@go build -o $(BIN_DIR)/arrietty ./cmd/arrietty/main.go

test:
	@go test ./tokenize
	@go test ./parse
	#@go test ./analyze
	@go test ./interpret

clean:
	@$(RM) $(BIN_DIR)/arrietty