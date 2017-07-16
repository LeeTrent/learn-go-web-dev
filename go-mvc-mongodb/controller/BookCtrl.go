package controller

import (
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/go-mvc-mongodb/model"
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/go-mvc-mongodb/view"
	"github.com/LeeTrent/ToddMcleod/learn-go-web-dev/go-mvc-mongodb/dao"
	"github.com/julienschmidt/httprouter"
	"database/sql"
	"net/http"
	"strconv"
	"errors"
	"fmt"
)

type BookController struct {
	bookDAO *dao.BookDAO
	bookView *view.BookView
}

func NewBookController() *BookController {

	fmt.Println("NewBookController()")

	return &BookController{ bookDAO: dao.NewBookDAO(), bookView: view.NewBookView() }
}

func (bc *BookController) CreateGet(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//fmt.Println("CreateGet()")

	if req.Method != http.MethodGet {
		http.Error(resp, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	bc.bookView.CreateGet(resp)
}

func (bc *BookController) CreatePost(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//fmt.Println("CreatePost()")

	if req.Method != http.MethodPost {
		http.Error(resp, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	book, err := bc.bookFromRequest(req)
	if err != nil {
		bc.bookView.Error(resp, err.Error())
		return
	}

	dbBook, err := bc.bookDAO.Create(book)
	if err != nil {
		//http.Error(resp, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		bc.bookView.Error(resp, err.Error())
		return
	}

	bc.bookView.CreatePost(resp, dbBook)
}

func (bc *BookController) UpdateGet(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {

	//fmt.Println("UpdateGet()")

	if req.Method != http.MethodGet {
		http.Error(resp, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	isbn := params.ByName("isbn")
	book, err := bc.bookDAO.RetrieveByISBN(isbn)

	if err != nil {
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Book not found for ISBN '%s'", isbn)
			bc.bookView.Error(resp, errMsg)
			return
		} else {
			//http.Error(resp, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			bc.bookView.Error(resp, err.Error())
			return
		}
	}

	bc.bookView.UpdateGet(resp, book)
}

func (bc *BookController) UpdatePost(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//fmt.Println("UpdatePost()")

	if req.Method != http.MethodPost {
		http.Error(resp, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	book, err := bc.bookFromRequest(req)
	if err != nil {
		bc.bookView.Error(resp, err.Error())
		return
	}

	dbBook, err := bc.bookDAO.Update(book)
	if err != nil {
		//http.Error(resp, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		bc.bookView.Error(resp, err.Error())
		return
	}

	bc.bookView.UpdatePost(resp, dbBook)
}

func (bc *BookController) DeleteGet(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {

	//fmt.Println("DeleteGet()")

	if req.Method != http.MethodGet {
		http.Error(resp, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	isbn := params.ByName("isbn")
	success, err := bc.bookDAO.Delete(isbn)

	if success != true {
		bc.bookView.Error(resp, err.Error())
		return
	}

	bc.RetieveAllGet(resp, req, params)
}

func (bc *BookController) RetieveAllGet(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//fmt.Println("RetieveAllGet()")

	if req.Method != http.MethodGet {
		http.Error(resp, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	books, err := bc.bookDAO.RetrieveAll()

	if err != nil {
		bc.bookView.Error(resp, err.Error())
		return
	}

	bc.bookView.RetrieveAllGet(resp, books)
}

func (bc *BookController) RetieveOneGet(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {

	//fmt.Println("RetieveOneGet()")

	if req.Method != http.MethodGet {
		http.Error(resp, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	isbn := params.ByName("isbn")
	book, err := bc.bookDAO.RetrieveByISBN(isbn)

	if err != nil {
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Book not found for ISBN '%s'", isbn)
			bc.bookView.Error(resp, errMsg)
			return
		} else {
			//http.Error(resp, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			bc.bookView.Error(resp, err.Error())
			return
		}
	}

	bc.bookView.RetrieveOneGet(resp, book)
}

func (bc *BookController) bookFromRequest(req *http.Request) (model.Book, error) {

	//fmt.Println("bookFromRequest()")

	book 		:= model.Book{}
	book.Isbn 	 = req.FormValue("isbn")
	book.Title 	 = req.FormValue("title")
	book.Author  = req.FormValue("author")
	strPrice 	:= req.FormValue("price")

	if book.Isbn == "" || book.Title == "" || book.Author == "" || strPrice == "" {
		return model.Book{}, errors.New("ISBN, Title, Author and Price are mandatory fields.")
	}

	f64Price, err := strconv.ParseFloat(strPrice, 32)
	if err != nil {
		return model.Book{}, err
	}
	book.Price = float32(f64Price)

	return book, nil
}