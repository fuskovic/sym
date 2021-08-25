.PHONY: fmt
fmt:
	@goimports -w *.go

.PHONY: lint
lint: fmt
	@golangci-lint run -v

.PHONY: test
test:
	@go test -v . 