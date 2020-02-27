package main

import (
	"encoding/json"
	"fmt"

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

	jsonBytes, _ := json.Marshal(exampleTodo)

	fmt.Println(string(jsonBytes))
}
