# Real Word golabg graphql sample

# Requirements
- Docker: https://www.docker.com/products/overview


# Get Start
- run ```docker-componse up mysql -d```
- change 192.168.0.106 to your internal ip in docker-compose.yml
- run ```docker-componse up web```

# Useage
Add Record
```
mutation {
 createUser(user: {name:"Russell", email:"russell@abc.com", password:"123456"})
  {
    name,
    email,
    updatedAt
  }
}
```
Query Record
```query {
  user(id:1) {
    name,
    email,
    updatedAt
  }
}```