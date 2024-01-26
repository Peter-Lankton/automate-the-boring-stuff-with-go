package main

import (
	"flag"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"strconv"
)

func main() {

	nPtr := flag.Int("number", 12, "how far to go on the table")
	savePtr := flag.String("save", "multiplication_table.xlsx", "for creating a multiplication table")
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

	style := xlsx.NewStyle()

	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.Border = border
	style.ApplyBorder = true

	fill := *xlsx.NewFill("solid", "00FF0000", "FF000000")
	style.Fill = fill
	style.ApplyFill = true

	font := *xlsx.NewFont(10, "Verdana")
	style.Font = font
	style.ApplyFont = true

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet")
	log.Printf("Sheet Name: %v", sheet.Name)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	for i := 1; i <= *nPtr; i++ {
		if i == 1 {
			row = sheet.AddRow()
			cell = row.AddCell()
		}

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetInt(i)
	}

	for i := 1; i <= *nPtr; i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetInt(i)
	}

	for i := 1; i <= *nPtr; i++ {
		for j := 1; j <= *nPtr; j++ {
			cell = sheet.Cell(i, j)
			cell.SetInt(i * j)
			log.Printf("times %v", strconv.Itoa(i)+" * "+strconv.Itoa(j))
		}
	}

	err = file.Save(pwdDir + "/" + *savePtr)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("save %v", *savePtr)

}
