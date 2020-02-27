package main

import (
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sebnyberg/proto-json-example/todo"
)

func main() {
	exampleTodo := &todo.Todo{
		TodoId:    uuid.New().String(),
		Status:    todo.Status_CREATED,
		Content:   "Do stuff",
		CreatedAt: ptypes.TimestampNow(),
	}

	m := jsonpb.Marshaler{}
	jsonStr, _ := m.MarshalToString(exampleTodo)

	fmt.Println(jsonStr)
}
