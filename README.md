# GO REST API

## Endpoints

| URL          | HTTP method | Auth | JSON Response     |
| ------------ | ----------- | ---- | ----------------- |
| /users/login | POST        |      | user's token      |
| /users       | GET         | Y    | all users         |
| /products    | GET         |      | all products      |
| /products    | POST        | Y    | new product added |
| /products    | PATCH       | Y    | edited product    |
| /products    | DELETE      | Y    | id                |

## Steps

1. Go extension for vscode

2. `go mod init github.com/rrd108/go-rest-api`

3. `go get -u github.com/gin-gonic/gin` - download, **u**pdate and install the package

4. Create main.go

5. `go get -u github.com/go-sql-driver/mysql`

6. login

7. `go get github.com/githubnemo/CompileDaemon` and `go get github.com/githubnemo/CompileDaemon`
   So we can run it by `~/go/bin/CompileDaemon -command="./go-rest-api"`

8. Create `launch.json` for debugging

9. `go get -u gorm.io/gorm` and `go get -u gorm.io/driver/mysql`

10. `go get -u github.com/gin-contrib/cors`
