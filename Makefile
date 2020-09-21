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

##install: Build and install go application executable
install:
	go install -v ./...
##build: Creating a service binary
build:
	$(GOBUILD) -o $(BINARY_NAME) $(PATH_MAIN_FILE)
##clear: Go clear
clean:
	go clean
##env: Print useful environment variables to stdout
env:
	echo PACKAGES $(PACKAGES)
	echo FILES $(FILES)

##code-analysis: Checking through static code analyzers
code-analysis: lint vet gosec
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

##test: Running unit tests
test:$(FILES)
	$(GOTEST) -v ./...
##test-cover: Run test coverage and generate html report
test-cover:
	rm -fr coverage
	mkdir coverage
	go list -f '{{if gt (len .TestGoFiles) 0}}"go test -covermode count -coverprofile {{.Name}}.coverprofile -coverpkg ./... {{.ImportPath}}"{{end}}' ./... | xargs -I {} bash -c {}
	echo "mode: count" > coverage/cover.out
	grep -h -v "^mode:" *.coverprofile >> "coverage/cover.out"
	rm *.coverprofile
	go tool cover -html=coverage/cover.out -o=coverage/cover.html
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

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'