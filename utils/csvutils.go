package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCSVFile(fileToParse string) ([][]string, error) {
	csvfile, err := os.Open(fileToParse)
	if err != nil {
		fmt.Println(err)
		return [][]string{}, err
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()

	return rawCSVdata, err
}

func WriteCSVToFile(filePath string, contents string) {
	f, err := os.Create(filePath)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	f.WriteString(contents)
}
