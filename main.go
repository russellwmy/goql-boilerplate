package main

import (
	"./data"
	"./graph"

	"os"
	"fmt"
	"net/http"
	"github.com/graphql-go/handler"
	"golang.org/x/net/context"
)


func main() {

	
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	
	
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		connectionString := os.Getenv("CONNECTION_STRING")
		db := data.Connect(connectionString)
		ctx := context.WithValue(context.Background(), "db", db)
		
		graph.Init(db)
		h := handler.New(&handler.Config{
			Schema: &graph.Schema,
			Pretty: true,
		})
		h.ContextHandler(ctx, w, r)
	})

	fmt.Println("Now server is running on port 3000")
	fmt.Println("Get single todo: curl -g 'http://localhost:3000/graphql?query={todo(id:\"b\"){id,text,done}}'")
	fmt.Println("Create new todo: curl -g 'http://localhost:3000/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'")
	fmt.Println("Load todo list: curl -g 'http://localhost:3000/graphql?query={todoList{id,text,done}}'")
	fmt.Println(port)
	http.ListenAndServe(port, nil)
}
