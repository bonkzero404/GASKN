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
│       ├── permission.go
│       └── rate_limiter.go
├── casbin_rbac_model.conf
├── config
│   └── config.go
├── database
│   ├── driver
│   │   ├── casbin.go
│   │   ├── connector.go
│   │   ├── mysql.go
│   │   └── postgresql.go
│   ├── migration.go
│   └── stores
│       ├── client_assignment_store.go
│       ├── client_store.go
│       ├── role_client_store.go
│       ├── role_store.go
│       ├── role_user_client_store.go
│       ├── role_user_store.go
│       ├── user_activation_store.go
│       └── user_store.go
├── dto
│   ├── mail_dto.go
│   └── response_dto.go
├── go.mod
├── go.sum
├── lang
│   ├── en
│   │   ├── auth.json
│   │   ├── client.json
│   │   ├── global.json
│   │   ├── middleware.json
│   │   ├── role.json
│   │   └── user.json
│   └── id
│       ├── auth.json
│       ├── client.json
│       ├── global.json
│       ├── middleware.json
│       ├── role.json
│       └── user.json
├── logs
│   ├── access.log
│   └── sql.error.log
├── main.go
├── modules
│   ├── auth
│   │   ├── contracts
│   │   │   └── user_auth_service.go
│   │   ├── dto
│   │   │   ├── user_auth_profile_dto.go
│   │   │   ├── user_auth_request_dto.go
│   │   │   └── user_auth_response_dto.go
│   │   ├── handlers
│   │   │   └── auth_handler.go
│   │   ├── module.go
│   │   ├── route.go
│   │   └── services
│   │       └── auth_service.go
│   ├── client
│   │   ├── contracts
│   │   │   ├── client_repository.go
│   │   │   └── client_service.go
│   │   ├── dto
│   │   │   ├── client_request_dto.go
│   │   │   └── client_response_dto.go
│   │   ├── handlers
│   │   │   └── client_handler.go
│   │   ├── module.go
│   │   ├── repositories
│   │   │   └── client_repository.go
│   │   ├── route.go
│   │   └── services
│   │       └── client_service.go
│   ├── role
│   │   ├── contracts
│   │   │   ├── role_client_repository.go
│   │   │   ├── role_client_service.go
│   │   │   ├── role_repository.go
│   │   │   └── role_service.go
│   │   ├── dto
│   │   │   ├── role_request_dto.go
│   │   │   └── role_response_dto.go
│   │   ├── handlers
│   │   │   ├── role_client_handler.go
│   │   │   └── role_handler.go
│   │   ├── module.go
│   │   ├── repositories
│   │   │   ├── role_client_repository.go
│   │   │   └── role_repository.go
│   │   ├── route.go
│   │   └── services
│   │       ├── role_client_service.go
│   │       └── role_service.go
│   └── user
│       ├── contracts
│       │   ├── repository_aggregate.go
│       │   ├── user_activation_factory.go
│       │   ├── user_activation_repository.go
│       │   ├── user_forgot_pass_factory.go
│       │   ├── user_repository.go
│       │   └── user_service.go
│       ├── dto
│       │   ├── user_activation_request_dto.go
│       │   ├── user_create_reponse_dto.go
│       │   ├── user_create_request_dto.go
│       │   ├── user_forgot_pass_act_request_dto.go
│       │   ├── user_forgot_pass_request_dto.go
│       │   └── user_reactivation_request_dto.go
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
├── postman
│   └── NgWork.postman_collection.json
├── seeders
│   ├── casbin.go
│   ├── roles.go
│   ├── seed_abstraction.go
│   ├── seeds.go
│   ├── user.go
│   └── user_role.go
├── templates
│   ├── user_activation.html
│   └── user_forgot_password.html
└── utils
    ├── api_group.go
    ├── api_wrapper.go
    ├── casbin_adapter.go
    ├── create_token.go
    ├── file.go
    ├── hash.go
    ├── logs.go
    ├── mail.go
    ├── pagination.go
    ├── rand_char.go
    ├── route_features.go
    ├── translation.go
    └── validation_struct.go
```

# How to run this project?

To run this project copy the .env.example file into .env, then do the configuration as you need

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
