# TODO API
## Documentation
- Golang: https://go.dev/

- gRPC: https://grpc.io/docs/

- GORM: https://gorm.io/docs/index.html
## Requirement
- Docker
- grpcurl
## Architecture
The application follows a clean architecture pattern to keep the code simple, organizable, and maintainable. The main components are:
-  **Transport Layer**: Handles gRPC requests and responses.

-  **Business Layer**: Contains the business logic.

-  **Repository Layer**: Interacts with the database using GORM.
## Start Development
### Step 1: Start Containers
Start the server and database containers using Docker Compose:
```bash
docker-compose  --profile  dev  up
```
### Step 2: Migrate All Database
Build the production image and run the database migrations:
```bash
docker-compose  --profile  prod  build

docker-compose  run  --rm  api-prod  migrate
```
if you want to migrate a single table
```bash
docker-compose  run  --rm  api-prod  migrate  -table  table_name
```
## Sample gRPC Commands
### 1. Sign Up
```bash
grpcurl -d '{ "name": "Nguyen Van A", "password": "123456" }' -plaintext localhost:50051 user.TodoService/SignUp
```
### 2. Login
```bash
grpcurl -d '{ "name": "ducpv", "password": "123456" }' -plaintext localhost:50051 user.TodoService/Login
```
### 3. Create Task
```bash
grpcurl -H "Authorization: Bearer <your-token>" -d '{"user_id":"1", "title": "New Task", "description": "Task description", "progress": "do", "priority":"1"}' -plaintext localhost:50051 task.TodoService/CreateTask
```
Example:
```bash
grpcurl -H "Authorization: Bearer <eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjF9LCJleHAiOjE4MzQ5Mjc0NDMsImlhdCI6MTczNDkyNzQ0M30.N0OkuJzmP3_pNwk5urtKwSH8__vKwglKGbHmDTTeu6kJQDPSpoLYcmUKC8dXCOTnnTRwTYeBwCNwaLPRjKxeZw>" -d '{"user_id":"1", "title": "New Task", "description": "Task description", "progress": "do", "priority":"1"}' -plaintext localhost:50051 task.TodoService/CreateTask
```
### 4. Read Task
```bash
grpcurl -d '{"user_id":"1"}' -H 'Authorization: Bearer <token>' -plaintext localhost:50051 task.TodoService/ReadTask
```
### 5.  Update Task
```bash
grpcurl -d '{"id":3,"title":"Update Title"}' -H 'Authorization: Bearer <token>' -plaintext localhost:50051 task.TodoService/UpdateTask
```
### 6. Delete Task
```bash
grpcurl -d '{"id":6}' -H 'Authorization: Bearer <token>' -plaintext localhost:50051 task.TodoService/DeleteTask
```
## What I Love About My Solution
-   The clean architecture pattern makes the codebase easy to understand and maintain.
-   Using Docker ensures that the application can be easily run locally and in different environments.
-   gRPC provides a high-performance and scalable way to handle communication between services.
## What Else I Want You to Know
-   Due to time constraints, some advanced features and optimizations were not implemented.
-   Additional unit tests and integration tests would further improve the reliability of the application.
-   Future improvements could include better error handling and more comprehensive logging.