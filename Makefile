PACKAGES=$(shell go list ./... | grep -v /vendor/)

protoc: ## build proto files and inject gorm tags
	protoc -I . api/todo/v1/*.proto  --go_out=plugins=grpc,paths=source_relative:.
	protoc-go-inject-tag -input=api/todo/v1/todo.pb.go -XXX_skip=gorm
