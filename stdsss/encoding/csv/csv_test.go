package csv_test

import (
	"encoding/csv"
	"io"
	"os"
	"testing"
)

func TestCsvReader(t *testing.T) {
	path := "test.csv"
	f, _ := os.Open(path)
	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		t.Log(record)
	}
}

func TestCsvWriter(t *testing.T) {
	target := "target.csv"
	f, _ := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	csvWriter := csv.NewWriter(f)
	records := [][]string{
		{"Bob", "001", "002", "003", "004", "005", "006"},
		{"GoodMan", "001", "002", "003", "004", "005", "006"},
	}
	csvWriter.WriteAll(records)
}
