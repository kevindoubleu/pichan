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
	chmod +x out/habits.out
	out/habits.out -config=configs/config.yaml

# CGO_ENABLED=0; Enables statically linked binaries to make the application more portable
# the golang:alpine docker image does not have glibc binary which would've been be needed
build-habits:
	CGO_ENABLED=0 \
	GOOS=linux \
		go build \
			-o out/habits.out \
			cmd/habits/habits.go

init-configs:
	cp configs/config.yaml.example configs/config.yaml
	cp configs/test_config.yaml.example configs/test_config.yaml
