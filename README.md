# Go GraphQL MongoDB Project

Simple GraphQL backend using:
- Go
- GraphQL
- MongoDB
- gqlgen

## Features

- GraphQL API
- MongoDB connection
- Query users
- Ready for adding data later

---

## Setup

### 1. Install dependencies

```bash
go mod tidy
```

### 2. Generate gqlgen files

```bash
go run github.com/99designs/gqlgen generate
```

### 3. Run MongoDB

Make sure MongoDB is running locally on:

```
mongodb://localhost:27017
```

### 4. Start server

```bash
go run server.go
```

---

## GraphQL Playground

Open:

```
http://localhost:8080
```

---

## Example Query

```graphql
query {
  users {
    id
    name
    email
  }
}
```
### nested query
```txt
query{
  clients{
    id
    sites{
      id
      siteId
      gatewayInstances{
        id
      }
    }
  }
}
```
