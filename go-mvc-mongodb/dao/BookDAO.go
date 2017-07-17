package dao

import (
	"fmt"
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/go-mvc-mongodb/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbURL            = "mongodb://MyApplication:MyPassword@localhost/bookstore"
	dbName           = "bookstore"
	dbCollectionName = "books"
)

type BookDAO struct {
	books *mgo.Collection
}

func NewBookDAO() *BookDAO {

	fmt.Println("NewBookDAO()")

	// Get MongoDB session
	mgoSession, err := mgo.Dial(dbURL)

	if err != nil {
		panic(err)
	}

	fmt.Println("Obtained MongoDB Session")

	// Make sure we have a live session
	if err = mgoSession.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Pinged MongoDB Session")

	// Now that we have a live session,
	// get the database that we need
	// from the session ("bookstore")
	mgoDB := mgoSession.DB(dbName)

	fmt.Printf("Obtained the '%s' database\n", dbName)

	// Now that we have the database that we need,
	// get the collection that this DAO will be
	// using("books")
	mgoCollection := mgoDB.C(dbCollectionName)

	fmt.Printf("Obtained the '%s' collection\n", dbCollectionName)

	return &BookDAO{books: mgoCollection}
}

func (dao *BookDAO) Create(book model.Book) (model.Book, error) {

	err := dao.books.Insert(book)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (dao *BookDAO) Update(book model.Book) (model.Book, error) {

	err := dao.books.Update(bson.M{"isbn": book.Isbn}, &book)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (dao *BookDAO) Delete(isbn string) (bool, error) {

	err := dao.books.Remove(bson.M{"isbn": isbn})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (dao *BookDAO) RetrieveAll() ([]model.Book, error) {

	books := []model.Book{}

	err := dao.books.Find(bson.M{}).All(&books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (dao *BookDAO) RetrieveByISBN(isbn string) (model.Book, error) {

	book := model.Book{}

	err := dao.books.Find(bson.M{"isbn": isbn}).One(&book)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (dao *BookDAO) RetrieveByTitle(title string) ([]model.Book, error) {

	return dao.retrieve("title", title)
}

func (dao *BookDAO) RetrieveByAuthor(author string) ([]model.Book, error) {

	return dao.retrieve("author", author)
}

func (dao *BookDAO) retrieve(label, value string) ([]model.Book, error) {

	books := []model.Book{}

	err := dao.books.Find(bson.M{label: value}).All(&books)
	if err != nil {
		return nil, err
	}

	return books, nil
}
