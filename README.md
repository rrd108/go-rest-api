# GO REST API

## Endpoints

| URL          | HTTP method | Auth | JSON Response     |
| ------------ | ----------- | ---- | ----------------- |
| /users/login | POST        |      | user's token      |
| /users       | GET         | Y    | all users         |
| /products    | GET         |      | all products      |
| /products    | POST        | Y    | new product added |
| /products    | PATCH       | Y    | edited product    |
| /products    | DELETE      | Y    | true / false      |

## Steps

1. Go extension for vscode

2. Create `go/src/github.com/rrd108/go-rest-api`

3. `go mod init`

4. `go get -u github.com/gin-gonic/gin` - download, **u**pdate and install the package

5. Create main.go

6. `go get -u github.com/go-sql-driver/mysql`

7. login

8. `go get github.com/githubnemo/CompileDaemon` and `go get github.com/githubnemo/CompileDaemon`
   So we can run it by `~/go/bin/CompileDaemon -command="./go-rest-api"`

CORS
