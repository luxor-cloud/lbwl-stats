GO=$(shell which go)

test:
	$(GO) generate
	$(GO) test -cover ./...
build:
	$(GO) build -o statsserver main.go
