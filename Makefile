PACKAGES=$(shell go list ./... | grep -v /vendor/)

build: ## build the go packages
	@go build $(PACKAGES)
