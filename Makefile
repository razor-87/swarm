B=$(shell git rev-parse --abbrev-ref HEAD)
BRANCH=$(subst /,-,$(B))
GITREV=$(shell git describe --abbrev=7 --always --tags)
REV=$(GITREV)-$(BRANCH)
GORUN=go run
GOBUILD=GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

info:
	- @echo "revision $(REV)"

lint:
	@golangci-lint run

test-problems:
	go test --tags=problems -v .

test:
	go test -v ./...

test-race:
	go test -race -timeout=60s -count 1 ./...

run:
	@$(GORUN) -v .

run-race:
	@$(GORUN) -race .

escape: info
	@$(GOBUILD) -v -gcflags "-m -m" && rm -rf ./swarm

build:
	@$(GOBUILD) -ldflags "-s -w"

.PHONY: info lint test-problems test test-race run run-race escape build
