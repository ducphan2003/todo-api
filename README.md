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
### Step 1: Create file .env
```bash
PG_USER=postgres
PG_PASSWORD=Admin123
PG_DATABASE=todo
PG_PORT=5432
PG_HOST=todo-db
PG_CONNECT_TIMEOUT=300
PG_TIMEZOME=UTC
```
### Step 2: Start Containers
Start the server and database containers using Docker Compose:
```bash
docker-compose --profile dev build

docker-compose --profile dev up -d
```
### Step 3: Migrate All Database (please run dev)
`production`:Build the production image and run the database migrations:
```bash
docker-compose --profile prod build

docker-compose run --rm api-prod migrate
```
if you want to migrate a single table
```bash
docker-compose  run  --rm  api-prod  migrate  -table  table_name
```
`dev`:Build the production image and run the database migrations:
```bash
docker cp ./data_init.sql todo-db:/docker-entrypoint-initdb.d/data_init.sql

docker exec -i todo-db psql -U postgres -d todo -f /docker-entrypoint-initdb.d/data_init.sql
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
## What Else I Want You to Know about however i do not have enough time to complete
1. Bổ sung, tổ chức database
2. Xác thực người dùng có quyền thực hiện với data của họ
Cấu trúc transportation/grpc thành các file theo chức năng
Example:
transportation
|---grpcuser
    |---create.go
    |---update.go
    ...
3. Định nghĩa lại response của gRPC:
message CommonResponse {
  int32 code = 1;          // Mã trạng thái (0: thành công, >0: lỗi)
  string message = 2;      // Thông báo lỗi hoặc thành công
  google.protobuf.Any data = 3; // Dữ liệu kết quả (nếu thành công)
}
