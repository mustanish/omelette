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
func Exist(docKey string) (bool, error) {
	connectors.OpenCollection("tokens")
	_, err := connectors.ReadDocument(docKey, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
