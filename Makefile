GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -cover
GOGET=$(GOCMD) get
BINARY_NAME=neon-auth
BINARY_UNIX=$(BINARY_NAME)_unix


build:
	$(GOBUILD) -o bin/auth-neon src/cmd/main.go
test: 
	$(GOTEST) -v ./...
integration-test:
	$(GOTEST) -v ./... -tags=integration