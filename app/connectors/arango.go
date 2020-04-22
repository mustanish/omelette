package connectors

import (
	"context"
	"log"
	"os"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/mustanish/omelette/app/config"
)

var (
	db  driver.Database
	col driver.Collection
)

// Initialize creates/opens database
func Initialize(config *config.Config) {
	var (
		client driver.Client
		ctx    = context.Background()
	)
	if conn, err := http.NewConnection(http.ConnectionConfig{Endpoints: []string{os.Getenv("DATABASE_URL")}}); err != nil {
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
	}
	createIndexes()
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
	var ctx = driver.WithReturnNew(context.Background(), &document)
	meta, err := col.CreateDocument(ctx, document)
	if err != nil {
		log.Println("FAILED::could not create document because of", err.Error())
		return "", err
	}
	return meta.Key, nil
}

// UpdateDocument updates a single document in the collection based on key passed
func UpdateDocument(key string, document interface{}) (string, error) {
	var ctx = driver.WithReturnNew(context.Background(), &document)
	meta, err := col.UpdateDocument(ctx, key, &document)
	if err != nil {
		log.Println("FAILED::could not update document because of", err.Error())
		return "", err
	}
	return meta.Key, nil
}

// ReadDocument reads a single document from the collection based on key passed
func ReadDocument(key string, result interface{}) (string, error) {
	meta, err := col.ReadDocument(nil, key, nil)
	if err != nil {
		log.Println("FAILED::could not read document because of", err.Error())
		return "", err
	}
	return meta.Key, nil
}

// RemoveDocument remove a single document from the collection based on key passed
func RemoveDocument(key string) (string, error) {
	meta, err := col.RemoveDocument(nil, key)
	if err != nil {
		log.Println("FAILED::could not remove document because of", err.Error())
		return "", err
	}
	return meta.Key, nil
}

// QueryDocument runs query and returns a cursor to iterate over the returned document
func QueryDocument(query string, bindVars map[string]interface{}) driver.Cursor {
	if len(bindVars) == 0 {
		bindVars = nil
	}
	cursor, err := db.Query(nil, query, bindVars)
	if err != nil {
		log.Println("FAILED::could not query document because of", err.Error())
	}
	defer cursor.Close()
	return cursor
}

// Drop removes database
func Drop() {
	db.Remove(context.Background())
}

func createIndexes() {
	var (
		ctx = context.Background()
	)
	for key, value := range index {
		OpenCollection(key)
		for _, inValue := range value.([]map[string]interface{}) {
			switch inValue["type"] {
			case "ttl":
				var option driver.EnsureTTLIndexOptions
				option.Name = inValue["name"].(string)
				col.EnsureTTLIndex(ctx, inValue["field"].(string), inValue["expireAfter"].(int), &option)
				break
			case "persistent":
				var option driver.EnsurePersistentIndexOptions
				option.Name = inValue["name"].(string)
				option.Sparse = true
				option.Unique = true
				col.EnsurePersistentIndex(ctx, inValue["fields"].([]string), &option)
				break
			}
		}
	}
}
