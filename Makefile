GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -cover
GOGET=$(GOCMD) get
BINARY_NAME=neon-auth
BINARY_UNIX=$(BINARY_NAME)_unix
PROJECTNAME=$(shell basename "$(PWD)")

PACKAGES=$(shell go list ./...)
FILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

##build: Creating a service binary
build:
	$(GOBUILD) -o $(BINARY_NAME) $(PATH_MAIN_FILE)
##go clean
clean:
	go clean
##test: Running unit tests
test:$(FILES)
	$(GOTEST) -v ./...
##env: Print useful environment variables to stdout
env:
	echo PACKAGES $(PACKAGES)
	echo FILES $(FILES)
##lint:  run go lint on the source files
lint: 
	golint $(PACKAGES)


help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'