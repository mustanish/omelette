package connectors

import (
	"context"
	"log"
	"omelette/config"
	"os"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
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

// CreateDocument creates a single document in the collection
func CreateDocument(collection string, document interface{}) (string, error) {
	err := openCollection(collection)
	if err != nil {
		return "", err
	}
	ctx := driver.WithReturnNew(context.Background(), &document)
	meta, err := col.CreateDocument(ctx, document)
	if err != nil {
		log.Println("FAILED::could not create document because of", err.Error())
		return "", err
	}
	return meta.Key, nil
}

// UpdateDocument updates a single document in the collection based on key passed
func UpdateDocument(collection string, key string, document interface{}) (string, error) {
	err := openCollection(collection)
	if err != nil {
		return "", err
	}
	ctx := driver.WithReturnNew(context.Background(), &document)
	meta, err := col.UpdateDocument(ctx, key, document)
	if err != nil {
		log.Println("FAILED::could not update document because of", err.Error())
		return "", err
	}
	return meta.Key, nil
}

// ReadDocument reads a single document from the collection based on key passed
func ReadDocument(collection string, key string, result interface{}) (string, error) {
	err := openCollection(collection)
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	meta, err := col.ReadDocument(ctx, key, result)
	if err != nil {
		log.Println("FAILED::could not read document because of", err.Error())
		return "", err
	}
	return meta.Key, nil
}

// RemoveDocument remove a single document from the collection based on key passed
func RemoveDocument(collection string, key string) (string, error) {
	err := openCollection(collection)
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	meta, err := col.RemoveDocument(ctx, key)
	if err != nil {
		log.Println("FAILED::could not remove document because of", err.Error())
		return "", err
	}
	return meta.Key, nil
}

// QueryDocument runs query and returns a cursor to iterate over the returned document
func QueryDocument(query string, bindVars map[string]interface{}) (driver.Cursor, error) {
	ctx := driver.WithQueryCount(context.Background())
	cursor, err := db.Query(ctx, query, bindVars)
	defer cursor.Close()
	return cursor, err
}

// Drop removes database
func Drop() {
	db.Remove(context.Background())
}

func openCollection(name string) error {
	ctx := context.Background()
	exist, err := db.CollectionExists(ctx, name)
	if err != nil {
		log.Println("FAILED::could not check if collection exist because of ", err.Error())
		return err
	}
	if exist {
		col, err = db.Collection(ctx, name)
		if err != nil {
			log.Println("FAILED::could open collection because of ", err.Error())
			return err
		}
	} else {
		col, err = db.CreateCollection(ctx, name, nil)
		if err != nil {
			log.Println("FAILED::could not create collection because of ", err.Error())
			return err
		}
	}
	return nil
}

func createIndexes() {
	var (
		ctx = context.Background()
	)
	for key, value := range index {
		openCollection(key)
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
