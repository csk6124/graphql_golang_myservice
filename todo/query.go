package todo

import (
	"context"
	"myservice/graphql/todo/gen/graphqlmodel"
)

// NewTodo 서비스
func NewTodo(
	ctx context.Context,
) Todo {
	once.Do(func() {
		instance = &todo{
			ctx: ctx,
		}
	})
	return instance
}

// Todo 조회
func (s *todo) Todo(
	ctx context.Context,
	where graphqlmodel.TodoWhere,
) (*graphqlmodel.Todo, error) {
	return &graphqlmodel.Todo{
		Title:   "할일",
		Content: "내일",
		Type:    graphqlmodel.TodoTypeDayly,
	}, nil
}
