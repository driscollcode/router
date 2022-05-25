.PHONY: mocks
mocks:
	@cd mock; go generate

.PHONY: test
test: mocks
	@go test
