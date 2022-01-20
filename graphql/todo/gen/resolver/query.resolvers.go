package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myservice/graphql/todo/gen"
	"myservice/graphql/todo/gen/graphqlmodel"
	"myservice/todo"
	"time"
)

func (r *queryResolver) Todo(ctx context.Context, where graphqlmodel.TodoWhere) (*graphqlmodel.Todo, error) {
	ctx, done, errDone, cancel := TimeoutContextInit(
		ctx,
		time.Second*TimoutValue,
	)
	defer cancel()

	go func() {
		s := todo.NewTodo(ctx)
		data, err := s.Todo(ctx, where)
		if err != nil {
			errDone <- err
		}
		done <- data
	}()

	response, err := ResponseInterface(ctx, done, errDone)
	if err != nil {
		return nil, err
	}
	ret := response.(*graphqlmodel.Todo)
	return ret, nil
}

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
const TimoutValue = 5
