## Documentation

- Golang: https://go.dev/
- Gin: https://github.com/gin-gonic/gin
- GORM: https://gorm.io/docs/index.html

## Requirement

- Docker

# Start Development

## Step 1 -> start 2 containers: server and db

```bash
docker-compose --profile dev up
```

## Step 2 -> Migrate All Database

```bash
docker-compose --profile prod build
docker-compose run --rm api-prod migrate
```

_if you want to migrate single table_

```bash
docker-compose run --rm api-prod migrate -table table_name
```
grpcurl -d '{ "name": "abc", "password": "passwrd123" }' -plaintext localhost:50051 user.TodoService/SignUp

grpcurl -d '{ "name": "abc", "password": "passwrd123" }' -plaintext localhost:50051 user.TodoService/Login

grpcurl -H "Authorization: Bearer <your-token>" -d '{"user_id":"1", "title": "New Task", "description": "Task description", "progress": "do", "priority":"1"}' -plaintext localhost:50051 task.TodoService/CreateTask

grpcurl -H "Authorization: Bearer <eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjF9LCJleHAiOjE4MzQ5Mjc0NDMsImlhdCI6MTczNDkyNzQ0M30.N0OkuJzmP3_pNwk5urtKwSH8__vKwglKGbHmDTTeu6kJQDPSpoLYcmUKC8dXCOTnnTRwTYeBwCNwaLPRjKxeZw>" -d '{"user_id":"1", "title": "New Task", "description": "Task description", "progress": "do", "priority":"1"}' -plaintext localhost:50051 task.TodoService/CreateTask

grpcurl -d '{"user_id":"5"}' -plaintext localhost:50051 task.TodoService/ReadTask

grpcurl -d '{"id":3,"title":"so 3"}' -plaintext localhost:50051 task.TodoService/UpdateTask

grpcurl -d '{"id":6}' -plaintext localhost:50051 task.TodoService/DeleteTask