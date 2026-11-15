package main

import (
	"RESTAPI/http"
	"RESTAPI/todo"
	"fmt"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started on port 9091")
	}
}
