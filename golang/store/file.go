package store

import (
	"encoding/csv"
	"io"
	"os"
)

type StoreFile struct {
	Map map[string]Values
}

func (s *StoreFile) Get(alias string) Values {
	return s.Map[alias]
}

func FileStore(location string) (*StoreFile, error) {

	csvFile, err := os.Open(location)
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	m := make(map[string]Values)
	csvReader := csv.NewReader(csvFile)
	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		m[fields[0]] = Values{fields[1], fields[2], fields[3], fields[4], fields[5]}
	}
	return &StoreFile{Map: m}, nil

}
