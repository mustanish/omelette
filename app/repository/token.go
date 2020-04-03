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

// ReadToken reads token from database
func (t *Token) ReadToken(docKey string) error {
	connectors.OpenCollection("tokens")
	return nil
}
