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
