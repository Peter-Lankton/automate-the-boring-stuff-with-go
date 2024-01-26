package main

import (
	"flag"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"strconv"
)

var opts struct {
	FileOPEN string `short:"o" long:"open" default:"table.xlsx" description:"Файл таблицы"`
}

func main() {
	openPtr := flag.String("open", "table.xlsx", "file to open")
	flag.Parse()

	pwdDir, _ := os.Getwd()

	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	log.SetOutput(fLog)

	excelFileName := pwdDir + "/" + *openPtr

	log.Printf("File to open: %v", excelFileName)

	wb, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	i := wb.Sheets[0]
	log.Printf("Name: %v", i.Name)

	for m := 0; m < i.MaxCol; m++ {
		strName := "saveFile" + strconv.Itoa(m+1) + ".txt"
		saveFile, err := os.Create(strName)
		log.Printf("File name: %v", saveFile.Name())
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		defer saveFile.Close()

		for u := 0; u < i.MaxRow; u++ {
			sd := i.Cell(u, m)
			ss := sd.String()
			saveFile.WriteString(ss)
			saveFile.WriteString("\n")
		}
	}
}
