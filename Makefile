GO=$(shell which go)
DOCKER=$(shell which docker)

test:
	$(GO) generate
	$(GO) test -cover ./...
build:
	$(GO) build -o statsserver main.go
startpg:
	cd hack && ./run_db.sh
stoppg:
	$(DOCKER) stop statspg
