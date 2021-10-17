.PHONY: fmt
fmt:
	@goimports -w *.go

.PHONY: lint
lint:
	@golangci-lint run -v

.PHONY: test
test:
	@go clean -testcache && go test -v . 

.PHONY: coverage
coverage:
	@./scripts/generate_test_coverage_report.sh $(mode)

.PHONY: update_badge
update_badge:
	@gopherbadger -md="README.md" -png=false