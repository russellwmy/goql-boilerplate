# Requirements
- Docker: https://www.docker.com/products/overview


# Get Start
- run ```docker-componse up mysql -d```
- change 192.168.0.106 to your internal ip in docker-compose.yml
- run ```docker-componse up web```

# Useage

web_1    | Now server is running on port 3000
web_1    | Get single todo: curl -g 'http://localhost:3000/graphql?query={todo(id:"b"){id,text,done}}'
web_1    | Create new todo: curl -g 'http://localhost:3000/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'
web_1    | Load todo list: curl -g 'http://localhost:3000/graphql?query={todoList{id,text,done}}'
web_1    | :3000
