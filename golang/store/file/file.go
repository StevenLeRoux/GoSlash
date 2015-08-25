package file

import (
	//"encoding/csv"
	"github.com/boltdb/bolt"
	//"io"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
	//"os"

	"le-roux.info/goslash/golang/store/common"
)

type StoreFile struct {
	db *bolt.DB
}

func (s *StoreFile) Update() {}

func (s *StoreFile) Get(alias string) (common.Values, bool) {

	var value common.Values

	// retrieve the data
	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("goslash"))
		if bucket == nil {
			return fmt.Errorf("Bucket goslash not found!")
		}

		temp := bucket.Get([]byte(alias))
		if temp == nil {
			return errors.New("No value for key : " + alias)
		}

		json.Unmarshal(temp, &value)

		return nil
	})

	if err != nil {
		log.Println(err)
		return common.Values{}, false
	}

	return value, true
}

func (s *StoreFile) Put(v common.Values) error {

	// initialize  bucket
	err := s.db.Update(func(tx *bolt.Tx) error {
		log.Println("s.db.Update(fn)")
		bucket, err := tx.CreateBucketIfNotExists([]byte("goslash"))
		if err != nil {
			log.Fatal("file: Put(): tx.CreateBucketIfNotExists():", err)
			return err
		}
		j, err := json.Marshal(v)
		if err != nil {

			log.Fatal(err)
		}

		err = bucket.Put([]byte(v.Alias), j)

		if err != nil {
			log.Fatal("Put() - bucket.Put() ", err)
		}
		return nil
	})

	if err != nil {
		log.Fatal("file: Put(): db.Update():", err)
	}

	return nil
}

func FileStore(location string) (*StoreFile, error) {

	db, err := bolt.Open(location, 0644, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal("FileStore() - bolt.Open()", err)
	}
	//defer db.Close()

	return &StoreFile{db: db}, nil

}
