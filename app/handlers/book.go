package handlers

import (
	"log"
	"math"
	"net/http"
	"omelette/app/connectors"
	"omelette/app/constants"
	"omelette/app/responses"
	"omelette/app/schemas/model"
	"omelette/app/schemas/validation"
	"omelette/helpers"
	"strconv"
	"strings"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jinzhu/copier"
)

// AddBook is used to add a book
func AddBook(res http.ResponseWriter, req *http.Request) {
	var response responses.Book
	book := new(model.Book)
	userID := req.Context().Value("ID").(string)
	data := req.Context().Value("data").(*validation.AddBook)
	book.AddedBy = userID
	copier.Copy(&book, data)
	book.CreatedAt = time.Now().Unix()
	docKey, err := connectors.CreateDocument("books", book)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	response.ID = docKey
	copier.Copy(&response, book)
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// SingleBook is used get detial of a single book
func SingleBook(res http.ResponseWriter, req *http.Request) {
	var response responses.Book
	book := new(model.Book)
	ID := chi.URLParam(req, "id")
	docKey, err := connectors.ReadDocument("books", ID, book)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	response.ID = docKey
	copier.Copy(&response, book)
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}

// UpdateBook is used to update details of an existing book
func UpdateBook(res http.ResponseWriter, req *http.Request) {
	var response responses.Book
	book := new(model.Book)
	ID := chi.URLParam(req, "id")
	data := req.Context().Value("data").(*validation.UpdateBook)
	docKey, err := connectors.ReadDocument("books", ID, book)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	} else if len(docKey) == 0 {
		render.Render(res, req, responses.NewHTTPError(http.StatusBadRequest, constants.NotFoundResource))
		return
	}
	copier.Copy(&book, data)
	book.UpdatedAt = time.Now().Unix()
	_, err = connectors.UpdateDocument("books", ID, book)
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
		query      strings.Builder
		response   responses.AllBook
		books      []responses.Book
		totalCount float64
	)
	book := new(responses.Book)
	pagination := make(map[string]int)
	bindVars := make(map[string]interface{})
	page, _ := strconv.Atoi(helpers.QueryParam(req, "page"))
	perPage, _ := strconv.Atoi(helpers.QueryParam(req, "perPage"))
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	totalPage := 1.0

	query.WriteString("FOR b IN books LET count= LENGTH(books) SORT b.name ASC LIMIT @offset,@limit ")
	query.WriteString("RETURN {'name': b.name,'genre': b.genre,'authorName': b.authorName,'addedBy': b.addedBy,")
	query.WriteString("'price': b.price,'coverImage': b.coverImage,'totalCount': count,'id':b._key}")
	bindVars["limit"] = perPage
	bindVars["offset"] = (page - 1) * perPage
	cursor, err := connectors.QueryDocument(query.String(), bindVars)
	if err != nil {
		render.Render(res, req, responses.NewHTTPError(http.StatusServiceUnavailable, constants.Unavailable))
		return
	}
	for {
		_, err := cursor.ReadDocument(nil, &book)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Println("FAILED::could not read book document because of", err.Error())
		} else {
			totalCount = book.TotalCount
			book.TotalCount = 0
			books = append(books, *book)
		}
	}
	if int(totalCount) > perPage {
		totalPage = ((totalCount) / float64(perPage))
	}
	pagination["totalItems"] = int(totalCount)
	pagination["perPage"] = perPage
	pagination["currentPage"] = page
	pagination["totalPages"] = int(math.Ceil(totalPage))
	response.Books = books
	response.Pagination = pagination
	render.Render(res, req, responses.NewHTTPSucess(http.StatusOK, response))
}
