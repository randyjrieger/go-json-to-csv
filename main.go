package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertJSONToCSV(JSONContent string) error {
	jsonFile, err := os.Open(JSONContent)

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	var ranking []CoolCarScale
	if err := json.NewDecoder(jsonFile).Decode(&ranking); err != nil {
		return err
	}

	for _, r := range ranking {
		var csvRow []string
		csvRow = append(csvRow, r.Make, r.Model, strconv.Itoa(r.Coolness))
		fmt.Println(strings.Join(csvRow, ","))
	}

	return nil
}

func main() {
	if err := convertJSONToCSV("coolCars.json"); err != nil {
		log.Fatal(err)
	}
}
