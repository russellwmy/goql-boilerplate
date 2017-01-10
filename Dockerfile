FROM russellwmy/golang

MAINTAINER Russell Wong <russellwmy@gmail.com>

RUN go get -v github.com/graphql-go/graphql && \
    go get -v github.com/graphql-go/handler && \
    go get -v github.com/graphql-go/relay && \
    go get -v github.com/jmoiron/sqlx && \
    go get -v github.com/go-sql-driver/mysql
    
