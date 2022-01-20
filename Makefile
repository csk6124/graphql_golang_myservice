gen:
	go get -d github.com/99designs/gqlgen
	cd graphql/todo && gqlgen

start: 
	go mod tidy
	go run main.go
