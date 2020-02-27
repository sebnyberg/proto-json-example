.PHONY: proto

proto:
	@protoc -I=proto --go_out=todo proto/todo.proto