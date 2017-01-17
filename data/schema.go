package data

var Schema = `
    scalar Time
    schema {
        query: Query
        mutation: Mutation
    }

    type Query {
        user(id: Int!): User
    }

    type Mutation {
        createUser(user: UserInput): User
    }
   
   input UserInput {
        name: String!
        email: String!
        password: String!
    }

    interface User {
        id: Int!
        name: String!
        email: String!
        updatedAt: Time
    }
`