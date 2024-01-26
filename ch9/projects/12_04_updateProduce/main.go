package main

import (
	"flag"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

func main() {
	openFilePtr := flag.String("open", "produceSales.xlsx", "Excel file to open")
	flag.Parse()
	pwdDir, _ := os.Getwd()

	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	log.SetOutput(fLog)

	excelFileName := pwdDir + "/" + *openFilePtr

	log.Printf("Excel file opened: %v", *openFilePtr)

	wb, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("error opening file%v", err)
	}

	sheet := wb.Sheet["Sheet"]

	priceUpdates := map[string]float64{
		"Lemon":  4.99,
		"Celery": 10.99,
		"Garlic": 7.47,
	}

	for rowNum := 1; rowNum < sheet.MaxRow; rowNum++ {
		produceName := sheet.Cell(rowNum, 0).Value
		if _, ok := priceUpdates[produceName]; ok {
			sheet.Cell(rowNum, 1).SetFloat(priceUpdates[produceName])
			log.Printf("Row: %v Produce: %v", rowNum, produceName)
		}
	}

	wb.Save("update" + *openFilePtr)

}
