build:
	@echo "==> Building package..."
	go build

unit-test:
	@echo "==> Running tests..."
	go test -v -parallel=4 -race -cover ./...
