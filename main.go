package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/wongpatrick/time-tracking-tool/internal/model"
	"github.com/wongpatrick/time-tracking-tool/internal/service"
)

func main() {
	fileContent, err := os.ReadFile("dataset/dataset.json")
	if err != nil {
		log.Fatal(err)
	}

	var shifts []model.Shift

	err = json.Unmarshal(fileContent, &shifts)
	if err != nil {
		log.Fatal(err)
	}

	employeeSummaries := service.ProcessShifts(shifts)

	outputJSON, err := json.MarshalIndent(employeeSummaries, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("dataset/output.json", outputJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Output written to dataset/output.json")
}
