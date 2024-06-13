SHELL := /bin/bash

dev:
	@./scripts/go_version.sh
	@echo "Going to run go mod tidy verify"
	@go mod tidy
	@go mod verify
	@go mod vendor
	go run main.go

update-all-packages:
	@echo "Warning: Going to update all packages"
	@echo "Going to run go get -v all / tidy / verify"
	@go get -v all
	@go mod tidy
	@go mod verify
	@go mod vendor

run-api:
	@./scripts/go_version.sh
	go run main.go

run-tests:
	@./scripts/go_version.sh
	go test ./... -v -coverpkg=./...
