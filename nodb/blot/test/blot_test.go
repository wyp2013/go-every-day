package test

import (
	"log"
	"testing"
)
import "github.com/boltdb/bolt"

func TestBlotDb(t *testing.T) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bucket, err := tx.CreateBucket([]byte("MyBucket"))
	if err != nil {
		return
	}
	bucket.Put([]byte("foo"), []byte("bar"))

}
