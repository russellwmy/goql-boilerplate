package graph

import (
    "github.com/graphql-go/graphql"
)



var todoType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Todo",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.Int,
        },
        "text": &graphql.Field{
            Type: graphql.String,
        },
        "done": &graphql.Field{
            Type: graphql.Boolean,
        },
    },
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
    Name: "RootMutation",
    Fields: graphql.Fields{
        /*
           curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'
        */
        "createTodo": &graphql.Field{
            Type: todoType, // the return type for this field
            Args: graphql.FieldConfigArgument{
                "text": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.String),
                },
            },
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {

                // marshall and cast the argument value
                text, _ := params.Args["text"].(string)

                // figure out new id

                // perform mutation operation here
                // for e.g. create a Todo and save to DB.
                newTodo := Todo{
                    Text: text,
                    Done: false,
                }

                CreateTodo (newTodo)

                // return the new Todo object that we supposedly save to DB
                // Note here that
                // - we are returning a `Todo` struct instance here
                // - we previously specified the return Type to be `todoType`
                // - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
                return newTodo, nil
            },
        },
    },
})

// root query
// we just define a trivial example here, since root query is required.
// Test with curl
// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
    Name: "RootQuery",
    Fields: graphql.Fields{

        /*
           curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'
        */
        "todo": &graphql.Field{
            Type: todoType,
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.String,
                },
            },
            Resolve: func(params graphql.ResolveParams) (interface{}, error) {

                idQuery, isOK := params.Args["id"].(int)
                if isOK {
                    // Search for el with id
                    todo := FindOne(idQuery)
                    return todo, nil
                }

                return Todo{}, nil
            },
        },

        /*
           curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'
        */
        "todoList": &graphql.Field{
            Type:        graphql.NewList(todoType),
            Description: "List of todos",
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return FindAll(), nil
            },
        },
    },
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query:    rootQuery,
    Mutation: rootMutation,
})