GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=neon-auth
BINARY_UNIX=$(BINARY_NAME)_unix
PROJECTNAME=$(shell basename "$(PWD)")

PACKAGES=$(shell go list ./...)
FILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")


all: 
	make clean
	make build
	make test

##install: Build and install go application executable
install:
	go install -v ./...
##build: Creating a service binary
build: generate-grpc 
	$(GOBUILD) -o bin/$(BINARY_NAME) cmd/neon-auth/main.go
	$(GOBUILD) -o bin/neon-migrate cmd/migrate/main.go
##clear: Go clear
clean:
	go clean -modcache
##env: Print useful environment variables to stdout
env:
	echo PACKAGES $(PACKAGES)
	echo FILES $(FILES)

##code-analysis: Checking through static code analyzers
code-analysis: lint vet gosec ineffassign misspell
##lint: run go lint on the source files
lint: 
	golint $(PACKAGES)
##vet: Go vet
vet:
	go vet $(PACKAGES)
##fmt: Format the go source files
fmt: 
	go fmt ./...
	goimports -w $(FILES)
##gosec: Check secure errors
gosec:
	gosec ./...
##ineffassign: Detect ineffectual assignments in Go code.
ineffassign:
	ineffassign ./*
##misspell: Correct commonly misspelled English words... quickly.
misspell:
	misspell ./*

##test: Running unit tests
test:$(FILES)
	$(GOTEST) -v ./...
##test-cover: Run test coverage and generate html report
test-cover:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...
##integration-test: Run integration tests
integration-test:$(FILES)
	$(GOTEST) -v ./... -tags=integration

##generate-grpc: Generate grpc files
generate-grpc:
	echo $(PWD)
	protoc -I=api/ --go_out=plugins=grpc:internal/interface/rpc  api/neon_auth.proto
##install-tools: Install vet,gosec,lint
install-tools:
	go get github.com/securego/gosec/v2/cmd/gosec
	go get -u golang.org/x/lint/golint
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/gordonklaus/ineffassign
	go get -u github.com/client9/misspell/cmd/misspell
	

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'