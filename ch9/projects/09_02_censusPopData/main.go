package main

import (
	"encoding/json"
	"flag"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

func main() {
	openExcelPtr := flag.String("open", "census_pop_data.xlsx", "for opening excel files.")
	saveToJson := flag.String("save", "data_to.json", "json file to save the excel file data to.")
	pwdDir, _ := os.Getwd()
	flag.Parse()

	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	log.SetOutput(fLog)

	exitData, err := os.OpenFile(pwdDir+`/`+*saveToJson, os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer exitData.Close()

	excelFileName := pwdDir + "/" + *openExcelPtr

	log.Printf("excel file name: %v", excelFileName)

	wb, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}

	sheet := wb.Sheet["Population by Census Tract"]
	log.Printf("Sheet name:  %v", sheet.Name)

	k := sheet.MaxRow

	type Key struct {
		State, Country string
	}
	popData := make(map[Key]float64)
	tractsData := make(map[Key]int)

	for u := 1; u < k; u++ {
		state := sheet.Cell(u, 1)
		country := sheet.Cell(u, 2)
		pop := sheet.Cell(u, 3)
		nameStateStr := state.String()
		countryStr := country.String()

		popFl, err := pop.Float()
		if err != nil {
			log.Fatalf("error %v", err)
		}

		popData[Key{nameStateStr, countryStr}] += popFl

		tractsData[Key{nameStateStr, countryStr}]++
	}

	type SaveData struct {
		Name   Key     `json:"state and country name"`
		Pop    float64 `json:"populations"`
		Tracts int     `json:"tracts"`
	}

	log.Printf("saving json JSON %s", exitData.Name())

	newStr := "\n"
	for keys := range popData {
		saveData2D := &SaveData{
			Name:   keys,
			Pop:    popData[keys],
			Tracts: tractsData[keys]}
		saveData2B, _ := json.Marshal(saveData2D)
		saveData2B = append(saveData2B, newStr...)

		if _, err := exitData.Write(saveData2B); err != nil {
			log.Panicf("error savying json  %v", err)
		}
	}
	log.Println("done ")
}
