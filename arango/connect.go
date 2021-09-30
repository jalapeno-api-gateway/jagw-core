package arango

import (
	"context"
	"log"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type ArangoDbConfig struct {
	Server string
	User string
	Password string
	DbName string
}

var Config ArangoDbConfig

func InitializeArangoDbAdapter(config ArangoDbConfig) {
	Config = config
}

func queryArangoDbDatabase(ctx context.Context, query string) driver.Cursor {
	db := openArangoDbDatabase(ctx)
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		log.Fatalf("Could not create Cursor , %v", err)
	}
	defer cursor.Close()
	return cursor
}

func openArangoDbDatabase(ctx context.Context) driver.Database {
	arangoDbClient := connectToArangoDb()
	db, err := arangoDbClient.Database(ctx, Config.DbName)
	if err != nil {
		log.Fatalf("Could not open database, %v", err)
	}
	return db
}

func connectToArangoDb() driver.Client {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{Config.Server},
	})
	if err != nil {
		log.Fatalf("Could not connect to ArangoDb, %v", err)
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(Config.User, Config.Password),
	})
	if err != nil {
		log.Fatalf("Could not create new ArangoDb client, %v", err)
	}
	return c
}
