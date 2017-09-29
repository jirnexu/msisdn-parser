package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

var countries map[string]string
var operators map[string]string

func main() {
	err := loadData()
	if err != nil {
		log.Fatalln(err)
	}
}

func loadData() error {
	d, err := readFile("data/countries.csv")
	if err != nil {
		return err
	}
	countries = make(map[string]string, len(d))
	for _, v := range d {
		prefix := v[0]
		countryCode := v[1]
		countries[prefix] = countryCode
	}

	d, err = readFile("data/operators.csv")
	if err != nil {
		return err
	}
	operators = make(map[string]string, len(d))
	for _, v := range d {
		prefix := v[1]
		operator := v[2]
		operators[prefix] = operator
	}
	return nil
}

func readFile(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bufio.NewReader(f))
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}
