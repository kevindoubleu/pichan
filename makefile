test: lint
	go test -v ./...

lint:
	gofmt -w .
	goimports -w .
	go vet ./...
