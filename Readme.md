#customers-ms

```
/users-service/
├── cmd/
│   └── users/
│       └── main.go
├── internal/
│   ├── domain/
│   │   ├── entity/
│   │   │   └── user.go
│   │   ├── value_objects/
│   │   │   └── email.go
│   │   ├── events/
│   │   │   └── user_registered.go
│   │   └── interfaces/
│   │       └── auth_service.go  // Interface no domínio
│   ├── application/
│   │   ├── commands/
│   │   │   ├── register_user/
│   │   │   │   ├── register_user.go       // Command
│   │   │   │   └── register_user_handler.go // Command Handler
│   │   │   ├── update_profile/
│   │   │   │   ├── update_profile.go
│   │   │   │   └── update_profile_handler.go
│   │   │   ├── change_password/
│   │   │   │   ├── change_password.go
│   │   │   │   └── change_password_handler.go
│   │   ├── queries/
│   │   │   ├── get_user/
│   │   │   │   ├── get_user.go            // Query
│   │   │   │   └── get_user_handler.go    // Query Handler
│   │   │   ├── list_users/
│   │   │   │   ├── list_users.go
│   │   │   │   └── list_users_handler.go
│   ├── infrastructure/
│   │   ├── repositories/
│   │   │   └── user_repository.go
│   │   ├── auth/
│   │   │   └── jwt_auth.go  // Implementação na infra
│   │   ├── cache/
│   │   │   └── redis_cache.go
│   │   └── database/
│   │       └── postgres.go
│   └── interfaces/
│       ├── http/
│       │   ├── handlers/
│       │   │   ├── user_handler.go
│       │   │   └── auth_handler.go
│       │   └── middlewares/
│       │       ├── auth_middleware.go
│       │       └── rate_limiter.go
│       └── grpc/
│           └── user_service.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── logging/
│   │   └── logger.go
│   └── utils/
│       └── utils.go
└── go.mod
```
