# Chapter 9. Excelling with GO

## XLSX GO

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	pwdDir, _ := os.Getwd()
	os.Mkdir(pwdDir+"/log", 0740)
	fLog, err := os.OpenFile(pwdDir+"/log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	check(err, fLog)

	excelFileName := pwdDir + "/example.xlsx"

	xlFile, err := xlsx.OpenFile(excelFileName)
	check(err, fLog)

	i := xlFile.Sheets[0]
	fmt.Println(i.Name)

	y := i.Cell(0, 0)
	a := y.String()
	fmt.Println(a)

	k := i.MaxRow
	for u := 0; u < k; u++ {
		sd := i.Cell(u, 1)
		ss := sd.String()
		fmt.Println(ss)
	}
}

// err check to log
func check(err error, fLog *os.File) {
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(io.MultiWriter(fLog, os.Stdout))
}

```

## Census Population Data

```go
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

```

## Save Excel

```go
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

```

## Update Produce

```go
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

```

## Multiplication Table 

```go
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

```

## txt To XLSX

```go
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

```

## XLSX to TXT

```go
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

```

## Conclusion

