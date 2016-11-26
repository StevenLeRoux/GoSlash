package file

import (
	//"encoding/csv"
	"github.com/StevenLeRoux/goslash/golang/store/model"
	"github.com/boltdb/bolt"
	"github.com/spf13/viper"
	//"io"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
	//"os"
)

type StoreFile struct {
	db *bolt.DB
}

func (s *StoreFile) Reload() {}

func (s *StoreFile) Close() {
	s.db.Close()
}

func (s *StoreFile) Get(alias string) (model.Values, bool) {
	var value model.Values

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
		return model.Values{}, false
	}

	return value, true
}

func (s *StoreFile) Dump() ([]model.Values, bool) {
	var value model.Values
	values := []model.Values{}

	err := s.db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte("goslash"))
		if bucket == nil {
			return fmt.Errorf("Bucket goslash not found!")
		}

		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if v == nil {
				break
			}
			json.Unmarshal(v, &value)
			values = append(values, value)
		}
		return nil
	})

	if err != nil {
		log.Println(err)
		return []model.Values{}, false
	}

	return values, true
}

func (s *StoreFile) Put(v model.Values) error {
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

func NewFileStore(v *viper.Viper) (*StoreFile, error) {
	db, err := bolt.Open(v.GetString("store.location"), 0644, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal("FileStore() - bolt.Open()", err)
	}

	return &StoreFile{db: db}, nil

}
