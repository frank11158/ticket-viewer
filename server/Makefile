.PHONY: install deploy build login push clean tool lint help
.DEFAULT: help

ENV_LIST=$(shell basename -a -s .env env/*.env)

help:
	@echo "make install ENV=<env filename>: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
	@echo "make test: run all unit tests"

install:
	@go build -v .
ifneq ($(findstring $(ENV),$(ENV_LIST)),)
	cp ./env/$(ENV).env .env
endif

tool:
	go vet -composites=false ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -f server
	rm -rf docs
	rm -f .env
	go clean -i .

test: install
	go test -v -covermode=count -coverprofile=test.out ./...
	go tool cover -html=test.out