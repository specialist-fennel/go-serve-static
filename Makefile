build:
	@go build -o /bin/go
run: build
	@./bin/go