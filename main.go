package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wongpatrick/time-tracking-tool/internal/model"
	"github.com/wongpatrick/time-tracking-tool/internal/service"
)

func main() {
	fileContent, err := os.ReadFile("dataset/dataset.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var shifts []model.Shift

	err = json.Unmarshal(fileContent, &shifts)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	employeeSummaries := service.ProcessShifts(shifts)

	outputJSON, err := json.MarshalIndent(employeeSummaries, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	os.WriteFile("dataset/output.json", outputJSON, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Output written to dataset/output.json")
}
