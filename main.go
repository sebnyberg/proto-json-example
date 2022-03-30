package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sebnyberg/proto-json-example/todo"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	exampleTodo := &todo.Todo{
		TodoId:            uuid.New().String(),
		Status:            todo.Status_CREATED,
		Content:           "Do stuff",
		CreatedAt:         ptypes.TimestampNow(),
		ExternalViewCount: 0,
	}

	m := protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
		Multiline:       true,
		Indent:          "  ",
	}
	jsonBytes, err := m.Marshal(exampleTodo)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", jsonBytes)
}
