package arango

import (
	"log"

	"github.com/arangodb/go-driver"
)


func readDocument(document driver.DocumentMeta, err error) bool {
	if driver.IsNoMoreDocuments(err) {
		return false
	}
	if err != nil {
		log.Fatalf("Error while reading from ArangoDB, %v", err)
	}
	return true
}
