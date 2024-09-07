build:
	@go build -o bin/budgetMannager cmd/main.go

run: build
	@./bin/budgetMannager cmd/main.go

test:
	@go test -v ./...
