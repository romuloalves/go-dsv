test:
	@go test -cover $(go list ./... | grep -v /vendor/)
.PHONY: test