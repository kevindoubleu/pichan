test: lint
	mkdir -p coverage
	go test -v ./... -covermode=count -coverprofile coverage/coverage.out
	go tool cover -func coverage/coverage.out
	go tool cover -html coverage/coverage.out -o coverage/coverage.html

lint:
	gofmt -w .
	goimports -w .
	go vet ./...

open-coverage:
	open coverage/coverage.html

run-habits: build-habits
	go run cmd/habits/habits.go

build-habits:
	go build -o out/habits.out cmd/habits/habits.go
