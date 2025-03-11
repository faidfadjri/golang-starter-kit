# <strong>Project Guidance</strong>

### Folder Structure

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
│── .env
│── go.mod
│── go.sum
│── log.txt
│── README.MD
______________________________________________________________________

Flow :
Route -> Handler -> Use Case -> Repositories -> Database -> Response
```
