build:
	@go build -o bin/firstproject cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/firstproject