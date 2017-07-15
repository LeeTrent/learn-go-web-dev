package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/go-mvc-rdbms/model"
	"fmt"
)

const (
	insertSQL = "INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)"
	updateSQL = "UPDATE books SET isbn = $1, title=$2, author=$3, price=$4 WHERE isbn=$1;"
	deleteSQL = "DELETE FROM books WHERE isbn=$1;"
	retrieveAllSQL = "SELECT isbn, title, author, price FROM books"
	retrieveByIsbnSql = "SELECT isbn, title, author, price FROM books WHERE isbn = $1"
	retrieveByTitleSql = "SELECT isbn, title, author, price FROM books WHERE title like %$1%"
	retrieveByAuthorSql = "SELECT isbn, title, author, price FROM books WHERE author like %$1%"
)

type BookDAO struct {
	db *sql.DB
}

func NewBookDAO() *BookDAO {

	fmt.Println("NewBookDAO()")

	pdb, err := sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")

	if (err != nil) {
		panic(err)
	}

	fmt.Println("DB Connection attempt was successful")

	return &BookDAO{db: pdb}
}

func (dao *BookDAO) Create(book model.Book) (model.Book, error){

	_, err := dao.db.Exec(insertSQL, book.Isbn, book.Title, book.Author, book.Price)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (dao *BookDAO) Update(book model.Book) (model.Book, error) {

	_, err := dao.db.Exec(updateSQL, book.Isbn, book.Title, book.Author, book.Price)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (dao *BookDAO) Delete(isbn string) (bool, error) {

	_, err := dao.db.Exec(deleteSQL, isbn)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (dao *BookDAO) RetrieveAll() ([]model.Book, error) {

	rows, err := dao.db.Query(retrieveAllSQL)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]model.Book, 0)

	for rows.Next() {
		book := model.Book{}
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (dao *BookDAO) RetrieveByISBN(isbn string) (model.Book, error) {

	book := model.Book{}

	row := dao.db.QueryRow(retrieveByIsbnSql, isbn)

	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (dao *BookDAO) RetrieveByTitle(title string) ([]model.Book, error) {

	return dao.retrieve(retrieveByTitleSql, title)
}

func (dao *BookDAO) RetrieveByAuthor(author string) ([]model.Book, error) {

	return dao.retrieve(retrieveByAuthorSql, author)
}

func (dao *BookDAO) retrieve(sql string, filter string) ([]model.Book, error) {

	rows, err := dao.db.Query(sql, filter)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]model.Book, 0)

	for rows.Next() {
		book := model.Book{}
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

//func (dao *BookDAO) RetrieveAll() ([]model.Book, error) {
//
//	rows, err := dao.db.Query(retrieveAllSQL)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	books := make([]model.Book, 0)
//
//	for rows.Next() {
//		book := model.Book{}
//		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
//		if err != nil {
//			return nil, err
//		}
//		books = append(books, book)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return books, nil
//}