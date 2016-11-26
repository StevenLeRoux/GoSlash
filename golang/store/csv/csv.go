package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/StevenLeRoux/goslash/golang/store/model"
	"github.com/spf13/viper"
)

type CsvStruct struct {
	Map    map[string]model.Values
	config *viper.Viper
	file   *os.File
}

func (c *CsvStruct) Get(alias string) (model.Values, bool) {
	return c.Map[alias], true
}

func (c *CsvStruct) Dump() ([]model.Values, bool) {
	s := []model.Values{}
	for _, v := range c.Map {
		s = append(s, v)
	}
	return s, true
}

func (c *CsvStruct) Put(v model.Values) error {
	log.Println("Put():", v.Alias)
	c.Map[v.Alias] = model.Values{Alias: v.Alias, Target: v.Target, User: v.User, Created: v.Created, Modified: v.Modified, Description: v.Description}
	c.FlushCSV()
	return nil
}

func (c *CsvStruct) FlushCSV() bool {
	log.Println("FlushCSV():")
	c.Reload()
	csvfile, err := os.OpenFile(c.config.GetString("store.location"), os.O_RDWR, 0644)
	//defer csvfile.Close()
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
		log.Println("FlushCSV(): Writing : ", v.Alias, " to ", c.config.GetString("store.location"))
		err := writer.Write([]string{v.Alias, v.Target, v.User, v.Created, v.Modified, v.Description})
		if err != nil {
			log.Println("Error:", err)
			return false
		}
	}

	writer.Flush()
	return true

}

func (c *CsvStruct) Reload() {
	log.Println("Reload():")
	csvFile, err := os.Open(c.config.GetString("store.location"))

	defer csvFile.Close()
	if err != nil {
		panic(err)
	}

	//c.Map = make(map[string]model.Values)
	csvReader := csv.NewReader(csvFile)
	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		c.Map[fields[0]] = model.Values{fields[0], fields[1], fields[2], fields[3], fields[4], fields[5]}
	}

}

func (c *CsvStruct) Close() {
	c.file.Close()
}

func NewCsvStore(v *viper.Viper) (*CsvStruct, error) {
	csvFile, err := os.Open(v.GetString("store.location"))
	//	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	m := make(map[string]model.Values)
	csvReader := csv.NewReader(csvFile)
	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		m[fields[0]] = model.Values{fields[0], fields[1], fields[2], fields[3], fields[4], fields[5]}
	}

	return &CsvStruct{Map: m, config: v, file: csvFile}, nil

}
