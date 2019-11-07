GO=go
GO_V=$(GO) version
GO_B=$(GO) build
NAME=app
all:build
	
build:
	@echo "test build"
	@echo "--------version------------"
	$(GO_V)
	@echo ""
	@echo "buildding"
	$(GO_B) -o $(NAME) -v
	@echo "finish"