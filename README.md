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

1. Create `go/src/go-rest-api`
1. `go mod init go-rest`
1. `go get -u github.com/gin-gonic/gin` - download, **u**pdate and install the package
1. Create main.go
1. `go get -u github.com/go-sql-driver/mysql`
1. login

CORS
