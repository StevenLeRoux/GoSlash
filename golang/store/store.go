package store

import (
	"fmt"
	"github.com/StevenLeRoux/goslash/golang/store/csv"
	"log"

	"github.com/StevenLeRoux/goslash/golang/store/common"
	"github.com/StevenLeRoux/goslash/golang/store/file"
)

type Store interface {
	Get(string) (common.Values, bool)
	Put(common.Values) error
	Update()
}

func New(engine string, location string) (*Store, error) {
	var s Store
	var err error
	switch engine {
	case "csv":
		s, err = csv.CsvStore(location)
		if err != nil {
			log.Panicln(err)
		}
	case "file":
		s, err = file.FileStore(location)
		if err != nil {
			log.Panic(err)
		}
		//case "mongo":
		//	MongoStore()
		//case "redis":
		//	RedisStore()
	default:
		return nil, fmt.Errorf("Engine: %s can not match any available engine", engine)
	}
	return &s, nil
}
