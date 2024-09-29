format:
	gofmt -w .

test:
	go test -v ./...

linter:
	golangci-lint run ./...