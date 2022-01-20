// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphqlmodel

import (
	"fmt"
	"io"
	"strconv"
)

type Todo struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Type    TodoType `json:"type"`
}

type TodoWhere struct {
	TodoType TodoType `json:"todoType"`
}

type TodoType string

const (
	TodoTypeDayly  TodoType = "DAYLY"
	TodoTypeWeekly TodoType = "WEEKLY"
)

var AllTodoType = []TodoType{
	TodoTypeDayly,
	TodoTypeWeekly,
}

func (e TodoType) IsValid() bool {
	switch e {
	case TodoTypeDayly, TodoTypeWeekly:
		return true
	}
	return false
}

func (e TodoType) String() string {
	return string(e)
}

func (e *TodoType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TodoType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TodoType", str)
	}
	return nil
}

func (e TodoType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
