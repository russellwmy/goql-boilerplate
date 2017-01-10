package main

import (
	"./data"
	"./graph"

	"os"
	"fmt"
	"net/http"
	"github.com/graphql-go/handler"
)

func main() {

	connectionString := os.Getenv("CONNECTION_STRING")
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	
	data.ConnectDB(connectionString)
	graph.Init()
	h := handler.New(&handler.Config{
		Schema: &graph.Schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	fmt.Println("Now server is running on port 3000")
	fmt.Println("Get single todo: curl -g 'http://localhost:3000/graphql?query={todo(id:\"b\"){id,text,done}}'")
	fmt.Println("Create new todo: curl -g 'http://localhost:3000/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'")
	fmt.Println("Load todo list: curl -g 'http://localhost:3000/graphql?query={todoList{id,text,done}}'")
	fmt.Println(port)
	http.ListenAndServe(port, nil)
}
