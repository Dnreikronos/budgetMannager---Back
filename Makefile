build:
	@go build -o bin/budgetMannager

run: build
	@./bin/budgetMannager

test: 
	@go test -v ./...
