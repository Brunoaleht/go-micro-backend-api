build: 
	@echo "Building..."
	go build -o $(BINARY_NAME) $(CMD_DIR)/main.go

run: build
	@echo "Running..."
	./$(BINARY_NAME)

dev:
	@echo "Running in dev mode..."
	swag build run