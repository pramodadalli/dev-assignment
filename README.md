# dev-assignment
A RESTful API built with [Beego](https://beego.me/) and Go (Go 1.20.1) for file upload and management with JWT authentication and user-specific storage quotas.

### 1. Clone the Repo
HTTPS -- https://github.com/pramodadalli/dev-assignment.git 

or
SSH -- git@github.com:pramodadalli/dev-assignment.git

## Install Dependencies
go mod tidy

##  Run the Server
bee run
## Server will run at:
 http://localhost:8080

##  Features

- JWT-based user authentication
- File upload per user (user-specific folders)
- File storage quota enforcement
- List uploaded files
- Check remaining storage

## Folder Structure
beego-dev-assignment/ 
beego-file-manager/
├── conf/
│   └── app.conf
├── controllers/
│   └── user.go
│   └── file.go
├── middlewares/
│   └── jwt.go
├── models/
│   └── user.go
│   └── file.go
├── routers/
│   └── router.go
├── storage/
│   └── (user folders will be created here)
├── main.go
