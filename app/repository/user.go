package repository

import (
	"github.com/mustanish/omelette/app/connectors"
	"github.com/mustanish/omelette/app/models"
)

// User is top level instance
type User models.User

// Authenticate responsible for interaction with database
func (u *User) Authenticate() (string, error) {
	var docKey string
	connectors.OpenCollection("users")
	docKey, err := connectors.CreateDocument(*u)
	if err != nil {
		return docKey, err
	}
	return docKey, nil
}

// VerifyUser responsible for interaction with database
func (u *User) VerifyUser() (bool, error) {
	return true, nil
}

// UpdateUser responsible for interaction with database
func (u *User) UpdateUser(docKey string) (*User, string, error) {
	connectors.OpenCollection("users")
	docKey, err := connectors.UpdateDocument(docKey, u)
	if err != nil {
		return nil, docKey, err
	}
	return u, docKey, nil
}

// DeleteUser responsible for interaction with database
func (u *User) DeleteUser() (bool, error) {
	return true, nil
}

// GetUser responsible for interaction with database
func (u *User) GetUser(identity string) *User {
	return u
}

// Exist responsible for interaction with database
func (u *User) Exist(identity string) (*User, string, error) {
	var (
		docKey   string
		query    = "FOR u IN users FILTER u.email == @identity || u.phone == @identity || u._key == @identity return u"
		bindVars = map[string]interface{}{"identity": identity}
	)
	connectors.OpenCollection("users")
	docKey, err := connectors.QueryDocument(query, bindVars, u)
	if err != nil {
		return nil, docKey, err
	}
	return u, docKey, nil
}
