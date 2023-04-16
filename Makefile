tidy:
	gofmt -s -w ./

test:
	go test ./...
