package store

import (
	"fmt"
	"log"

	"github.com/StevenLeRoux/goslash/golang/store/csv"
	"github.com/StevenLeRoux/goslash/golang/store/file"
	"github.com/StevenLeRoux/goslash/golang/store/model"
	"github.com/spf13/viper"
)

// Store interface abstract the storage engine that can be either file (BoltDB),
// CSV for testing purpose or a database (not implemented yet)
type Store interface {
	Get(string) (model.Values, bool)
	Put(model.Values) error
	Reload()
	Close()
	Dump() ([]model.Values, bool)
}

// New() returns an instance of a Store interface backed by the provided engine type.
func New(v *viper.Viper) (*Store, error) {
	var s Store
	var err error
	switch v.GetString("store.type") {
	case "csv":
		s, err = csv.NewCsvStore(v)
		if err != nil {
			log.Panicln(err)
		}
	case "file":
		s, err = file.NewFileStore(v)
		if err != nil {
			log.Panic(err)
		}
	default:
		return nil, fmt.Errorf("Engine: %s can not match any available engine", v.GetString("store.type"))
	}
	return &s, nil
}
