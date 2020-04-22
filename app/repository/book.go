package repository

import (
	"log"

	"github.com/arangodb/go-driver"
	"github.com/mustanish/omelette/app/connectors"
	"github.com/mustanish/omelette/app/models"
	"github.com/mustanish/omelette/app/responses"
)

// Book represents an instance of book model
type Book models.Book

// AddBook responsible for interaction with database
func (b *Book) AddBook() (string, error) {
	connectors.OpenCollection("books")
	docKey, err := connectors.CreateDocument(b)
	if err != nil {
		return docKey, err
	}
	return docKey, nil
}

// UpdateBook responsible for interaction with database
func (b *Book) UpdateBook(docKey string) (string, error) {
	connectors.OpenCollection("books")
	docKey, err := connectors.UpdateDocument(docKey, b)
	if err != nil {
		return docKey, err
	}
	return docKey, nil
}

// DeleteBook responsible for interaction with database
func (b *Book) DeleteBook() {

}

// AllBooks responsible for interaction with database
func (b *Book) AllBooks(page int, perPage int) ([]responses.Book, float64, error) {
	var (
		err        error
		book       responses.Book
		books      []responses.Book
		limit      = perPage
		offset     = (page - 1) * perPage
		totalCount float64
		query      = "FOR b IN books LET count= LENGTH(books) SORT b.name ASC LIMIT @offset,@limit RETURN {'name': b.name," +
			"'genre': b.genre,'authorName': b.authorName,'addedBy': b.addedBy,'price': b.price,'coverImage': b.coverImage," +
			"'totalCount': count,'id':b._key}"
		bindVars = map[string]interface{}{"limit": limit, "offset": offset}
	)
	connectors.OpenCollection("books")
	cursor := connectors.QueryDocument(query, bindVars)
	for {
		_, err := cursor.ReadDocument(nil, &book)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Println("FAILED::could not read book document because of", err.Error())
		} else {
			totalCount = book.TotalCount
			book.TotalCount = 0
			books = append(books, book)
		}
	}
	return books, totalCount, err
}

// SingleBook responsible for interaction with database
func (b *Book) SingleBook(docKey string) (responses.Book, error) {
	var (
		err      error
		book     responses.Book
		query    = "FOR b IN books FILTER b._key == @key RETURN b"
		bindVars = map[string]interface{}{"key": docKey}
	)
	connectors.OpenCollection("books")
	cursor := connectors.QueryDocument(query, bindVars)
	for {
		meta, err := cursor.ReadDocument(nil, &book)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Println("FAILED::could not read book document because of", err.Error())
		} else {
			book.ID = meta.Key
		}
	}
	return book, err
}

// Exist responsible for interaction with database
func (b *Book) Exist(docKey string) (string, error) {
	connectors.OpenCollection("books")
	docKey, err := connectors.ReadDocument(docKey, b)
	if err != nil {
		return docKey, err
	}
	return docKey, nil
}
