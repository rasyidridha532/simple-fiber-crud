# Simple Fiber CRUD Clean Architecture

This is my simple backend CRUD with Golang with Clean Architecture. I create this project to explore myself with coding, especially with Golang. 

## Feature
- Web Framework: Fiber
- Database: MongoDB
- Validation: Go-Playground

## On Going Feature
- Logging: Zap (In Progress)
- Dependency Injection (currently manual)
- Unit Testing
- Username & Password Auth
- JWT Auth

## Configuration
| ENV         | Description                               |
|-------------|:------------------------------------------|
| GO_ENV      | app environment (development, production) |
| DB_NAME     | mongodb database name                     |
| DB_USER     | mongodb user                              |
| DB_PASSWORD | mongodb password                          |
| DB_HOST     | mongodb host                              |
| DB_PORT     | mongodb port                              |
| PORT        | app port                                  |

## Endpoint
- /dashboard
    - `GET` show monitoring from Fiber
- /api/v1/accounts
    - `GET` get all users
- /api/v1/account
    - `GET` get one user
    - `POST` create one user
    - `DELETE` delete one user

# Reference
[golang-clean-architecture by Eko Kurniawan Khannedy](https://github.com/khannedy/golang-clean-architecture)