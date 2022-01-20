PROJECT_PATH = $(PWD)
.PHONY: install all gen

start: 
	go mod tidy
	go run main.go

gen:
	go get -d github.com/99designs/gqlgen
	cd graphql/todo && gqlgen