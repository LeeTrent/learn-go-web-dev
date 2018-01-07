# Purpose
This application:
1. was created for pedagogical purposes
2. attempts to apply the Model-View-Controller (MVC)design pattern in Go.
3. attempts to conform to REST API principles by using the "julienschmidt/httprouter" library
4. provides CRUD operations on one collection named 'books' in a document database (MongoDB)

# Database
MongoDB running on localhost

## Database Setup:
### Start MongoDb at command line:
```
mongod
```
### Create 'bookstore' databse:
```
use bookstore
```
### Create and populate 'books' collection:
```
db.books.insert([{"isbn":"978-1505255607","title":"The Time Machine","author":"H. G. Wells","price":5.99},{"isbn":"978-1503261960","title":"Wind Sand \u0026 Stars","author":"Antoine","price":14.99},{"isbn":"978-1503261961","title":"West With The Night","author":"Beryl Markham","price":14.99}])
```
### Create Authorized DB User:
```
db.createUser(
  {
    user: "my-db-user",
    pwd: "my-db-pw",
    roles: [ { role: "readWrite", db: "bookstore" } ]
  }
)
```
### Exit mongo & restart with auth enabled
```
mongod --auth
```
```
mongo -u "my-db-user" -p "my-db-pw" --authenticationDatabase "bookstore"
```
### Database Setup Test:

```
use bookstore
```

```
show collections
```

```
db.books.find()
```

```
db.books.insert({"isbn" : "978-1503261777", "title" : "Never Say Never", "author" : "Ian Fleming", "price" : 24.99 })
```

```
db.books.find()
```

# How to run Go application:
1. Execute the following in root folder where main.go resides: ```go run *.go```
2. Paste the following URL in your favorite web browser: ```http://localhost:8080/```