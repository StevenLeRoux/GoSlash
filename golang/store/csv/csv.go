package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"le-roux.info/goslash/golang/store/common"
)

type CsvStruct struct {
	Map      map[string]common.Values
	location string
}

func (c *CsvStruct) Get(alias string) (common.Values, bool) {
	return c.Map[alias], true
}

func (c *CsvStruct) Put(v common.Values) error {
	log.Println("Put():", v.Alias)
	c.Map[v.Alias] = common.Values{Alias: v.Alias, Target: v.Target, User: v.User, Created: v.Created, Modified: v.Modified, Description: v.Description}
	c.FlushCSV()
	return nil
}

func (c *CsvStruct) FlushCSV() bool {
	log.Println("FlushCSV():")
	c.Update()
	csvfile, err := os.OpenFile(c.location, os.O_RDWR, 0644)
	defer csvfile.Close()
	if err != nil {
		log.Println("Error:", err)
		return false
	}

	writer := csv.NewWriter(csvfile)
	for k, v := range c.Map {
		if k != v.Alias {
			log.Panic("database inconsistency : ", k, " == ", v.Alias)
			return false
		}
		log.Println("FlushCSV(): Writing : ", v.Alias, " to ", c.location)
		err := writer.Write([]string{v.Alias, v.Target, v.User, v.Created, v.Modified, v.Description})
		if err != nil {
			log.Println("Error:", err)
			return false
		}
	}

	writer.Flush()
	return true

}

func (c *CsvStruct) Update() {
	log.Println("Update():")
	csvFile, err := os.Open(c.location)

	defer csvFile.Close()
	if err != nil {
		panic(err)
	}

	//c.Map = make(map[string]common.Values)
	csvReader := csv.NewReader(csvFile)
	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		c.Map[fields[0]] = common.Values{fields[0], fields[1], fields[2], fields[3], fields[4], fields[5]}
	}

}

func CsvStore(location string) (*CsvStruct, error) {

	csvFile, err := os.Open(location)
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	m := make(map[string]common.Values)
	csvReader := csv.NewReader(csvFile)
	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		m[fields[0]] = common.Values{fields[0], fields[1], fields[2], fields[3], fields[4], fields[5]}
	}

	return &CsvStruct{Map: m, location: location}, nil

}
