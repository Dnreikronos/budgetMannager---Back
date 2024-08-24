build:
	@go build -o bin/budgetMannager/cmd

run: build
	@./bin/budgetMannager/cmd

test:
	@go test -v ./...
