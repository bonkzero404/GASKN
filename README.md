# GoLang go-starterkit-project

File structure in this project.

```
go-starterkit-project
├── Makefile
├── README.md
├── app
│   ├── bootstrap.go
│   └── middleware
│       ├── authenticate.go
│       └── rate_limiter.go
├── config
│   └── config.go
├── database
│   ├── driver
│   │   ├── connector.go
│   │   ├── mysql.go
│   │   └── postgresql.go
│   └── migration.go
├── domain
│   ├── dto
│   │   ├── mail_dto.go
│   │   └── response_dto.go
│   └── stores
│       ├── user_activation.go
│       └── user_store.go
├── go.mod
├── go.sum
├── main.go
├── modules
│   ├── auth
│   │   ├── domain
│   │   │   ├── dto
│   │   │   │   ├── user_auth_profile_dto.go
│   │   │   │   ├── user_auth_request_dto.go
│   │   │   │   ├── user_auth_response_dto.go
│   │   │   │   └── user_auth_validation_dto.go
│   │   │   └── interfaces
│   │   │       └── user_auth_service_interface.go
│   │   ├── handlers
│   │   │   └── auth_handler.go
│   │   ├── module.go
│   │   ├── route.go
│   │   └── services
│   │       └── auth_service.go
│   └── user
│       ├── domain
│       │   ├── dto
│       │   │   ├── user_activation_request_dto.go
│       │   │   ├── user_activation_request_validation_dto.go
│       │   │   ├── user_create_reponse_dto.go
│       │   │   ├── user_create_request_dto.go
│       │   │   ├── user_create_request_validation_dto.go
│       │   │   ├── user_forgot_pass_act_request_dto.go
│       │   │   ├── user_forgot_pass_act_validation_dto.go
│       │   │   ├── user_forgot_pass_request_dto.go
│       │   │   ├── user_forgot_pass_validation_dto.go
│       │   │   ├── user_reactivation_request_dto.go
│       │   │   └── user_reactivation_validation_dto.go
│       │   └── interfaces
│       │       ├── repository_aggregate_interface.go
│       │       ├── user_activation_factory_interface.go
│       │       ├── user_activation_repository_interface.go
│       │       ├── user_forgot_pass_factory_interface.go
│       │       ├── user_repository_interface.go
│       │       └── user_service_interface.go
│       ├── handlers
│       │   └── user_handler.go
│       ├── module.go
│       ├── repositories
│       │   ├── repository_aggregate.go
│       │   ├── user_activation_repository.go
│       │   └── user_repository.go
│       ├── route.go
│       └── services
│           ├── factories
│           │   ├── activation_factory.go
│           │   ├── user_activation_factory.go
│           │   └── user_forgot_pass_factory.go
│           └── user_service.go
├── rbac_model.conf
├── templates
│   ├── user_activation.html
│   └── user_forgot_password.html
└── utils
    ├── api_group.go
    ├── api_wrapper.go
    ├── create_token.go
    ├── hash.go
    ├── mail.go
    ├── rand_char.go
    └── validation_struct.go
```

# How to run this project?

To run this project copy the .env.example file into .env, then do the configuration as you need, here is the env file

```conf
# Application port
APP_PORT=3000
APP_HOST=localhost

# Endpoint
API_WRAP=api
API_VERSION=v1

# Database Connection
# mysql | pgsql
DB_DRIVER=mysql
DB_HOST=
DB_NAME=
DB_USER=
DB_PASSWORD=
DB_PORT=3306

# Database Pool
DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=100

#JWT
JWT_SECRET=supersecret

# Mail
MAIL_HOST=smtp.mailtrap.io
MAIL_PORT=2525
MAIL_USERNAME=null
MAIL_PASSWORD=null
MAIL_FROM=admin@example.com

# Storages
STORAGE_DIR=storages

```

After you create the configuration file, create a database in MySQL or PostgreSQL with the appropriate name in the configuration file above.

Run the command in the root directory

```
go run main.go
```

or if you use makefile run the following command

```
make watch
```

This command has a "hot reload" feature, but you will need the <b>reflect</b> library to run the command

# API Specifications

## Register User

```http
POST /api/v1/user/register HTTP/1.1
Host: 127.0.0.1:3000
Content-Type: application/json

{
    "full_name": "Jhon Doe",
    "email": "jhon@example.com",
    "phone": "17287817212",
    "password": "mylongpassword"
}
```

## Activation User

```http
POST /api/v1/user/activation HTTP/1.1
Host: 127.0.0.1:3000
Content-Type: application/json

{
    "email": "jhon@example.com",
    "code": "XHHuRNyX2Gq4C1LiIEkO32EbQoPBvQhF"
}
```

## Re-Send Activation Code

```http
POST /api/v1/user/activation/re-send HTTP/1.1
Host: 127.0.0.1:3000
Content-Type: application/json

{
    "email": "jhon@example.com"
}
```

## Request Forgot Password

```http
POST /api/v1/user/request-forgot-password HTTP/1.1
Host: 127.0.0.1:3000
Content-Type: application/json

{
    "email": "jhon@example.com"
}
```

## Forgot Password

```http
POST /api/v1/user/forgot-password HTTP/1.1
Host: 127.0.0.1:3000
Content-Type: application/json
Content-Length: 158

{
    "email": "jhon@example.com",
    "password": "mychangepassword",
    "repeat_password": "mychangepassword",
    "code": "u6BiYwbWRthBCa4r0HcUQjdcTaa70tyo"
}
```

## Authentication

```http
POST /api/v1/auth HTTP/1.1
Host: 127.0.0.1:3000
Content-Type: application/json

{
    "email": "jhon@example.com",
    "password": "mylongpassword"
}

```

## Get Profile

```http
GET /api/v1/auth/me HTTP/1.1
Host: 127.0.0.1:3000
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mzc1NzUwNDMsImlkIjoiNTk1ZWY0N2UtZThkOS00MjZjLThmNzItMjk2NjFiNjRlN2JlIn0.ChyYZB_DJofyZhN7BuPFT8NeX3AEyfKNbZp1YVba8Fw
```

## Refresh Token

```http
GET /api/v1/auth/refresh-token HTTP/1.1
Host: 127.0.0.1:3000
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mzc1NzUwNDMsImlkIjoiNTk1ZWY0N2UtZThkOS00MjZjLThmNzItMjk2NjFiNjRlN2JlIn0.ChyYZB_DJofyZhN7BuPFT8NeX3AEyfKNbZp1YVba8Fw
```
