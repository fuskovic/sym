.PHONY: fmt
fmt:
	@goimports -w *.go

.PHONY: lint
lint:
	@golangci-lint run -v

.PHONY: test
test:
	@go test -v . 