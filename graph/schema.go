package graph

import (
    "./todo"

    "github.com/graphql-go/graphql"
)

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
    Name: "RootMutation",
    Fields: graphql.Fields{
        "createTodo": &todo.CreateTodoField,
    },
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
    Name: "RootQuery",
    Fields: graphql.Fields{
        "todo": &todo.TodoField,
        "todoList": &todo.TodoListField,
    },
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query:    rootQuery,
    Mutation: rootMutation,
})