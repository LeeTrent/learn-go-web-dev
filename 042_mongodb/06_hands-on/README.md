# Run this code

1. Start your server

```
go run main.go
```

## CREATE a user
Enter this at the terminal
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
```
## RETRIEVE a user
Enter this at the terminal
```
curl http://localhost:8080/user/<enter-user-id-here>
```
# DELETE a user
Enter this at the terminal
```
curl -X DELETE http://localhost:8080/user/<enter-user-id-here>
```