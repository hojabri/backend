# backend
A REST web-service sample project written in Golang using go-fiber, GORM and PostgreSQL

## How to run
- Make sure you have Go installed ([download](https://golang.org/dl/)).
- Make sure you have installed docker and docker-compose tools (for running PostgreSQL database)  ([instruction](https://docs.docker.com/compose/install/)).
- Pull and setup project DB with docker-compose tools and docker-compose.yml file

```
> make db-up
```

Output:

```
Starting postgresql ... done

```

- Run the project:
```
> make run
```

Output:

```
 ┌───────────────────────────────────────────────────┐ 
 │                    Fiber v2.7.1                   │ 
 │               http://127.0.0.1:3000               │ 
 │       (bound on host 0.0.0.0 and port 3000)       │ 
 │                                                   │ 
 │ Handlers ............ 24  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID .............. 5699 │ 
 └───────────────────────────────────────────────────┘ 
```

- You can use postman (https://www.postman.com/) tool to call endpoints

## List of endpoints:
Note: replace ```:user_id``` with user's ID
```
show <hello> message:                       GET         http://127.0.0.1:3000

Get all users:                              GET         http://127.0.0.1:3000/user
Get single user information:                GET         http://127.0.0.1:3000/user/:user_id
Create new user:                            POST        http://127.0.0.1:3000/user
Update user:                                PUT         http://127.0.0.1:3000/user/:user_id
Delete user:                                DELETE      http://127.0.0.1:3000/user/:user_id

Get all companies:                          GET         http://127.0.0.1:3000/company
Get single company information:             GET         http://127.0.0.1:3000/company/:company_id
Create new company:                         POST        http://127.0.0.1:3000/company
Update company:                             PUT         http://127.0.0.1:3000/company/:company_id
Delete company:                             DELETE      http://127.0.0.1:3000/company/:company_id

Get all company categories:                 GET         http://127.0.0.1:3000/companyCategory
Get single company category information:    GET         http://127.0.0.1:3000/companyCategory/:companyCategory_id
Create new company category:                POST        http://127.0.0.1:3000/companyCategory
Update company category:                    PUT         http://127.0.0.1:3000/companyCategory/:companyCategory_id
Delete company category:                    DELETE      http://127.0.0.1:3000/companyCategory/:companyCategory_id

```
## Call endpoints with curl command:

### curl:
command line tool and library
for transferring data with URLs 
- Make sure you have curl tool installed in your computer (https://curl.se/)

### Samples:

### Get all users:
```
curl --location --request GET 'http://127.0.0.1:3000/user'
```

Output:
```json
{
  "code": 200,
  "body": [
    {
      "ID": 2,
      "CreatedAt": "0001-01-01T01:55:52+01:55",
      "UpdatedAt": "2021-04-16T15:03:45.89629+03:00",
      "DeletedAt": null,
      "first_name": "Omid",
      "last_name": "Hojabri",
      "email": "omid.hojabri@backend.com",
      "phone": "+90123456790"
    },
    {
      "ID": 4,
      "CreatedAt": "0001-01-01T01:55:52+01:55",
      "UpdatedAt": "2021-04-16T15:17:35.699638+03:00",
      "DeletedAt": null,
      "first_name": "Bill",
      "last_name": "Gates",
      "email": "bill.gates@microsoft.com",
      "company_id": 2,
      "company": {
        "ID": 2,
        "CreatedAt": "2021-04-16T15:17:05.418797+03:00",
        "UpdatedAt": "2021-04-16T15:17:05.418797+03:00",
        "DeletedAt": null,
        "name": "Microsoft",
        "category_id": 2,
        "category": null,
        "website": "http://microsoft.com"
      }
    }
  ],
  "title": "GetAllUsers",
  "message": "All users"
}
```
### Create user:
```
curl --location --request POST 'http://127.0.0.1:3000/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name" : "Bill",
    "last_name" : "Gates",
    "email" : "bill.gates@microsoft.com",
    "company_id" : 2
}'
```

Output:
```json
{
    "code": 200,
    "body": {
        "ID": 4,
        "CreatedAt": "2021-04-16T15:16:08.114701+03:00",
        "UpdatedAt": "2021-04-16T15:16:08.114701+03:00",
        "DeletedAt": null,
        "first_name": "Bill",
        "last_name": "Gates",
        "email": "bill.gates@microsoft.com",
        "company_id": 2,
        "company": {
          "ID": 2,
          "CreatedAt": "2021-04-16T15:17:05.418797+03:00",
          "UpdatedAt": "2021-04-16T15:17:05.418797+03:00",
          "DeletedAt": null,
          "name": "Microsoft",
          "category_id": 2,
          "category": null,
          "website": "http://microsoft.com"
        }      
    },
    "title": "OK",
    "message": "new user added successfully"
}
```
## Data models

User:

```go
type User struct {
	gorm.Model
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email,omitempty" gorm:"uniqueIndex"`
	Phone     string  `json:"phone,omitempty"`
	CompanyID *uint    `json:"company_id,omitempty"`
	Company   *Company `json:"company,omitempty"`
}
```

Company:

```go
type Company struct {
	gorm.Model
	Name         string          `json:"name" gorm:"uniqueIndex"`
	Description  string          `json:"description,omitempty"`
	CategoryID   *uint            `json:"category_id,omitempty"`
	Category     *CompanyCategory `json:"category"`
	Website      string          `json:"website,omitempty"`
}
```

Company Category:

```go
type CompanyCategory struct {
	gorm.Model
	Name string `json:"name" gorm:"uniqueIndex"`
}
```

## GORM
The fantastic ORM library for Golang aims to be developer friendly.
- More information (https://gorm.io/)
- Full-Featured ORM
- Associations (has one, has many, belongs to, many to many, polymorphism, single-table inheritance)
- Hooks (before/after create/save/update/delete/find)
- Eager loading with Preload, Joins
- Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
- Context, Prepared Statement Mode, DryRun Mode
- Batch Insert, FindInBatches, Find/Create with Map, CRUD with SQL Expr and Context Valuer
- SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery
- Composite Primary Key, Indexes, Constraints
- Auto Migrations
- Logger
- Extendable, flexible plugin API: Database Resolver (multiple databases, read/write splitting) / Prometheus…
- Every feature comes with tests
- Developer Friendly

## Go-Fiber
Fiber is a Go web framework built on top of Fasthttp, the fastest HTTP engine for Go. It's designed to ease things up for fast development with zero memory allocation and performance in mind.
More information (https://gofiber.io/)

## PostgerSQL
PostgreSQL is a powerful, open source object-relational database system with over 30 years of active development that has earned it a strong reputation for reliability, feature robustness, and performance.
More information (https://www.postgresql.org/)
