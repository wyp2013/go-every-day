package test

import (
	"log"
	"testing"

	"github.com/boltdb/bolt"
)

func TestBlotDb(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin(true)
	if err != nil {

	}
	defer tx.Rollback()
	tx.CreateBucket([]byte("MyBucket"))

	_, err = tx.CreateBucket([]byte("MyBucket1"))
	if err != nil {
	}

	if err := tx.Commit(); err != nil {
	}

}
