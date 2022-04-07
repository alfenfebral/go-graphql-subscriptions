# Go Graphql Subscriptions

## Publisher
```bash
go run publisher/main.go
```

## GraphQL Server
```bash
go run server.go
```

Go to URL http://localhost:8080/ and add below
```
subscription {
    payload
}
```