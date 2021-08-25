.PHONY: lint
lint:
	@goimports -w *.go

.PHONY: test
test:
	@go test -v . 