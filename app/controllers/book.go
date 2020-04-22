package controllers

import (
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/copier"
	"github.com/mustanish/omelette/app/constants"
	"github.com/mustanish/omelette/app/helpers"
	"github.com/mustanish/omelette/app/repository"
	"github.com/mustanish/omelette/app/responses"
	bookschemas "github.com/mustanish/omelette/app/schemas/book"
)

// AddBook is used to add a book
func AddBook(res http.ResponseWriter, req *http.Request) {
	var (
		book     repository.Book
		userID   = req.Context().Value("ID").(string)
		data     = req.Context().Value("data").(*bookschemas.AddBook)
		now      = time.Now()
		response responses.Book
	)
	book.AddedBy = userID
	copier.Copy(&book, data)
	book.CreatedAt, book.UpdatedAt = now.Unix(), now.Unix()
	docKey, err := book.AddBook()
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response.ID = docKey
	copier.Copy(&response, book)
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// UpdateBook is used to update details of an existing book
func UpdateBook(res http.ResponseWriter, req *http.Request) {
	var (
		book        repository.Book
		ID          = chi.URLParam(req, "id")
		data        = req.Context().Value("data").(*bookschemas.UpdateBook)
		now         = time.Now()
		response    responses.Book
		docKey, err = book.Exist(ID)
	)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	copier.Copy(&book, data)
	book.UpdatedAt = now.Unix()
	_, err = book.UpdateBook(ID)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response.ID = docKey
	copier.Copy(&response, book)
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// DeleteBook is used to delete details of an existing book
func DeleteBook(res http.ResponseWriter, req *http.Request) {

}

// AllBooks is used to fetch details of all books in paginated format
func AllBooks(res http.ResponseWriter, req *http.Request) {
	var (
		book       repository.Book
		response   responses.AllBook
		pagination = make(map[string]int)
		page, _    = strconv.Atoi(helpers.QueryParam(req, "page"))
		perPage, _ = strconv.Atoi(helpers.QueryParam(req, "perPage"))
		totalPage  = 1.0
	)
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	doc, totalCount, err := book.AllBooks(page, perPage)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	if int(totalCount) > perPage {
		totalPage = ((totalCount) / float64(perPage))
	}
	pagination["totalItems"] = int(totalCount)
	pagination["perPage"] = perPage
	pagination["currentPage"] = page
	pagination["totalPages"] = int(math.Ceil(totalPage))
	response.Books = doc
	response.Pagination = pagination
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// SingleBook is used get detial of a single book
func SingleBook(res http.ResponseWriter, req *http.Request) {
	var (
		book     repository.Book
		ID       = chi.URLParam(req, "id")
		doc, err = book.SingleBook(ID)
	)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(doc.ID) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, doc))
}
