package todo

import (
    "github.com/graphql-go/graphql"
    "github.com/jmoiron/sqlx"
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

var CreateTodoField = graphql.Field{
    Type: todoType, // the return type for this field
    Args: graphql.FieldConfigArgument{
        "text": &graphql.ArgumentConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
    },
    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        db := p.Context.Value("db").(*sqlx.DB)
        text, _ := p.Args["text"].(string)
        newTodo := Todo{
            Text: text,
            Done: false,
        }
        CreateTodo (db, newTodo)
        return newTodo, nil
    },
}

var TodoField =  graphql.Field {
    Type: todoType,
    Args: graphql.FieldConfigArgument{
        "id": &graphql.ArgumentConfig{
            Type: graphql.String,
        },
    },
    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        db := p.Context.Value("db").(*sqlx.DB)
        idQuery, isOK := p.Args["id"].(int)
        if isOK {
            // Search for el with id
            todo := FindOne(db, idQuery)
            return todo, nil
        }

        return Todo{}, nil
    },
}
var TodoListField =  graphql.Field {
    Type:        graphql.NewList(todoType),
    Description: "List of todos",
    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        db := p.Context.Value("db").(*sqlx.DB)
        return FindAll(db), nil
    },
}
