# Purpose
This application:
1. was created for pedagogical purposes
2. attempts to apply the Model-View-Controller (MVC)design pattern in Go.
3. Attempts to conform to REST API principles by using the "julienschmidt/httprouter" library
4. provides CRUD operations on one table named 'books' using a Relational Database (PostgreSQL)

# Database
PostgreSQL running on localhost

## Table Description:
```
        Table "public.books"
 Column |          Type          | Modifiers 
--------+------------------------+-----------
 isbn   | character(14)          | not null
 title  | character varying(255) | not null
 author | character varying(255) | not null
 price  | numeric(5,2)           | not null
Indexes:
    "books_pkey" PRIMARY KEY, btree (isbn)
```
# How to run
1. Execute the following in root folder where main.go resides: ```go run *.go```
2. Paste the following URL in your favorite web browser: ```http://localhost:8080/```
