package todo

import (
	"context"
	"myservice/graphql/todo/gen/graphqlmodel"
	"sync"
)

var (
	instance *todo = nil
	once     sync.Once
)

type todo struct {
	ctx context.Context
}

// Todo 인터페이스
type Todo interface {
	Todo(
		ctx context.Context,
		where graphqlmodel.TodoWhere,
	) (*graphqlmodel.Todo, error)
}
