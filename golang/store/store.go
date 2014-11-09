package store

import "fmt"

type Values struct {
	Target   string
	User     string
	Creation string
	Modified string
	Desc     string
}

type Store interface {
	Get(string) Values
}

func New(engine string, location string) (*Store, error) {
	var s Store
	var err error
	switch engine {
	case "file":
		s, err = FileStore(location)
		if err != nil {
			panic(err)
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
