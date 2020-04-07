package connectors

import (
	"context"
	"log"

	"github.com/arangodb/go-driver"
)

var (
	db  driver.Database
	col driver.Collection
)

func init() {
	//var config, _ = config.LoadConfig()
	//InitializeDB(config)
}

// InitializeDB initializes database
func InitializeDB() {
	log.Println("called")
	/*var (
		client driver.Client
		ctx    = context.Background()
	)
	if conn, err := http.NewConnection(http.ConnectionConfig{Endpoints: []string{config.Database.DBUrl}}); err != nil {
		log.Println("FAILED::could Not connect to database because of", err.Error())
	} else if client, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(config.Database.DBUser, config.Database.DBPassword),
	}); err != nil {
		log.Println("FAILED::could not create client because of", err.Error())
	}
	if exist, _ := client.DatabaseExists(ctx, config.Database.DBName); exist {
		db, _ = client.Database(ctx, config.Database.DBName)
	} else {
		db, _ = client.CreateDatabase(ctx, config.Database.DBName, nil)
	}*/
}

// OpenCollection opens a collection within the database
func OpenCollection(name string) {
	var ctx = context.Background()
	if exist, _ := db.CollectionExists(ctx, name); exist {
		col, _ = db.Collection(ctx, name)
	} else {
		col, _ = db.CreateCollection(ctx, name, nil)
	}
}

// CreateDocument creates a single document in the collection
func CreateDocument(document interface{}) (string, error) {
	var (
		ctx    = driver.WithReturnNew(context.Background(), &document)
		docKey string
	)
	meta, err := col.CreateDocument(ctx, document)
	if err != nil {
		log.Println("FAILED::could not create document because of", err.Error())
		return docKey, err
	}
	return meta.Key, nil
}

// UpdateDocument updates a single document in the collection based on key passed
func UpdateDocument(key string, document interface{}) (string, error) {
	var (
		ctx    = driver.WithReturnNew(context.Background(), &document)
		docKey string
	)
	meta, err := col.UpdateDocument(ctx, key, document)
	if err != nil {
		log.Println("FAILED::could not update document because of", err.Error())
		return docKey, err
	}
	return meta.Key, nil
}

// RemoveDocument remove a single document from the collection based on key passed
func RemoveDocument(key string) (string, error) {
	var ctx = context.Background()
	meta, err := col.RemoveDocument(ctx, key)
	if err != nil {
		log.Println("FAILED::could not create document because of", err.Error())
		return "", err
	}
	return meta.ID.String(), nil
}

// QueryDocument runs query and returns a cursor to iterate over the returned document
func QueryDocument(query string, bindVars map[string]interface{}, result interface{}) (string, error) {
	var (
		ctx    = context.Background()
		docKey string
	)

	if len(bindVars) == 0 {
		bindVars = nil
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		log.Println("FAILED::could not query document because of", err.Error())
		return docKey, err
	}
	defer cursor.Close()
	meta, err := cursor.ReadDocument(ctx, &result)
	if driver.IsNoMoreDocuments(err) {
		return docKey, nil
	} else if err != nil {
		log.Println("FAILED::could not read document because of", err.Error())
		return docKey, err
	}
	return meta.Key, nil
}
