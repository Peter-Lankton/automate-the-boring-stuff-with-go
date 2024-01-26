package main

import (
	"flag"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

func main() {

	saveFilePtr := flag.String("save", "files.xlsx", "")
	flag.Parse()

	pwdDir, _ := os.Getwd()

	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	log.SetOutput(fLog)

	var (
		file  *xlsx.File
		sheet *xlsx.Sheet
		row   *xlsx.Row
		cell  *xlsx.Cell
	)

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet")
	log.Printf("Added Sheet: %v", sheet)
	if err != nil {
		log.Printf("Err adding sheet: %v", err)
	}

	for i := 0; i < 50; i++ {
		row = sheet.AddRow()
		log.Printf("Adding row %v", row)
		row.AddCell().SetString("Col A")
		row.AddCell().SetString("Col B")
		row.AddCell().SetString("Col C")
		cell = row.AddCell()
		log.Printf("Added columns: %v", cell)
		cell.Value = "Loren ipsum"
	}
	err = file.Save(pwdDir + "/" + *saveFilePtr)
	if err != nil {
		log.Printf("Err saving: %v", err)
	}
	log.Printf("Name of file:%v", *saveFilePtr)

}
