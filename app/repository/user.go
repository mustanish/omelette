package repository

import (
	"log"

	"github.com/arangodb/go-driver"
	"github.com/mustanish/omelette/app/connectors"
	"github.com/mustanish/omelette/app/models"
)

// User represents an instance of user model
type User models.User

// Authenticate responsible for interaction with database
func (u *User) Authenticate() (string, error) {
	var docKey string
	connectors.OpenCollection("users")
	docKey, err := connectors.CreateDocument(u)
	if err != nil {
		return docKey, err
	}
	return docKey, nil
}

// UpdateUser responsible for interaction with database
func (u *User) UpdateUser(docKey string) (string, error) {
	connectors.OpenCollection("users")
	docKey, err := connectors.UpdateDocument(docKey, u)
	if err != nil {
		return docKey, err
	}
	return docKey, nil
}

// Exist responsible for interaction with database
func (u *User) Exist(identity string) (string, error) {
	var (
		docKey   string
		query    = "FOR u IN users FILTER u.email == @identity || u.phone == @identity || u._key == @identity return u"
		bindVars = map[string]interface{}{"identity": identity}
	)
	connectors.OpenCollection("users")
	cursor := connectors.QueryDocument(query, bindVars)
	for {
		meta, err := cursor.ReadDocument(nil, &u)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Println("FAILED::could not read user document because of", err.Error())
		} else {
			docKey = meta.Key
		}
	}
	return docKey, nil
}
