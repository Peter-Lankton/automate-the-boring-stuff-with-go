package main

import (
	"bufio"
	"flag"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"strings"
)

func main() {
	savePtr := flag.String("save", "table.xlsx", "convert txt file to an excel file")

	pwdDir, _ := os.Getwd()

	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	log.SetOutput(fLog)

	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Sheet")
	log.Printf("Sheet name %v", sheet.Name)
	if err != nil {
		log.Printf("Error adding sheet %v", err)
	}

	dir, err := os.Open(pwdDir)
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	dataFiles := make(map[string][]string)

	for _, fi := range fileInfos {
		if strings.HasSuffix(fi.Name(), ".txt") {
			log.Printf("File name: %v", fi.Name())
			fileOpen, err := os.Open(fi.Name())
			if err != nil {
				log.Fatalln(err)
			}
			var arr []string
			scanner := bufio.NewScanner(fileOpen)
			for scanner.Scan() {
				arr = append(arr, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			defer fileOpen.Close()

			dataFiles[fi.Name()] = arr
		}
	}

	i := 0
	for keys, prs := range dataFiles {
		sheet.Cell(0, i).SetValue(keys)
		l := 1
		for _, str := range prs {
			sheet.Cell(l, i).SetValue(str)
			l++
		}
		i++
	}

	err = file.Save(pwdDir + "/" + *savePtr)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Saving as...%v", *savePtr)
}
