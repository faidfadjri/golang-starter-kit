# <strong>README</strong>

## Table of Contents

- [Introduction](#introduction)
- [Structure](#structure)
- [Usage](#usage)

## Introduction

Akastra Mobile API is the core API of the Akastra Mobile App, built using Go (Golang) with several dependencies such as:

| Dependency                                   | Description                                    | Link                                       |
| -------------------------------------------- | ---------------------------------------------- | ------------------------------------------ |
| [go-chi](https://github.com/go-chi/chi)      | Lightweight, idiomatic web router for Go.      | [GitHub](https://github.com/go-chi/chi)    |
| [GORM](https://gorm.io/)                     | The ORM for database interactions in Go.       | [Website](https://gorm.io/)                |
| [Air](https://github.com/cosmtrek/air)       | Live reloading for Go applications.            | [GitHub](https://github.com/cosmtrek/air)  |
| [godotenv](https://github.com/joho/godotenv) | Load environment variables from a `.env` file. | [GitHub](https://github.com/joho/godotenv) |
| [MySQL](https://www.mysql.com/)              | The primary database used in the project.      | [Website](https://www.mysql.com/)          |

## Structure

```
______________________________________________________________________

📁akastra-mobile-api
│── 📁cmd
│   └── main.go
│── 📁src
│   ├── 📁app
│   │   ├── 📁entities
│   │   ├── 📁usecase
│   ├── 📁infrastructure
│   │   ├── 📁database
│   │   ├── 📁repositories
│   ├── 📁interface
│   │   ├── 📁handler
│   │   ├── 📁response
│   │   ├── 📁route
│   │   │
│── go.mod
│── go.sum
│── log.txt
│── README.MD
______________________________________________________________________

Flow :
Route -> Handler -> Use Case -> Repositories -> Database -> Response
```

## Usage

Let's get started to running this awesome project 😃

Copy `.env.example` and renamed it as `.env`

```
cp .env.example .env
```

> make sure to fill the required variable `DB_HOST`, `DB_NAME`, `DB_USER`, `DB_PASS`

Database migration using `gorm automigrate`

```
go run migration.go
```

Install `air-verse` (if it doesn't exist)

```bash
go get -u github.com/go-chi/chi/v5
```

Install required dependencies

```bash
go install
```

Running `air-verse`

```bash
air
```

Now, the awesome project is completely running 🎉
