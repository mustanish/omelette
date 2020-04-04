package repository

import (
	"github.com/mustanish/omelette/app/connectors"
	"github.com/mustanish/omelette/app/models"
)

// Token represents an instance of token model
type Token models.Token

// AddToken adds token into database
func (t *Token) AddToken() error {
	connectors.OpenCollection("tokens")
	_, err := connectors.CreateDocument(t)
	if err != nil {
		return err
	}
	return nil
}

// RemoveToken removes token from database
func (t *Token) RemoveToken(docKey string) error {
	connectors.OpenCollection("tokens")
	_, err := connectors.RemoveDocument(docKey)
	if err != nil {
		return err
	}
	return nil
}

// Exist checks if token exist in database or not
func Exist(key string) (bool, error) {
	var (
		docKey   string
		query    = "FOR t IN tokens FILTER t._key == @key return t"
		bindVars = map[string]interface{}{"key": key}
	)
	connectors.OpenCollection("tokens")
	docKey, err := connectors.QueryDocument(query, bindVars, nil)
	if err != nil {
		return false, err
	}
	if len(docKey) > 0 {
		return true, nil
	}
	return false, nil
}
