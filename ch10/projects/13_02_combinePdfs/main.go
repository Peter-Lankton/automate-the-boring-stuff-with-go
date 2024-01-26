package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
	pdfL "github.com/ledongthuc/pdf"
)

func main() {

	pwdDir, _ := os.Getwd()
	fLog, err := os.OpenFile(pwdDir+`/.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	defer fLog.Close()

	log.SetOutput(fLog)

	dir, err := os.Open(pwdDir)
	if err != nil {
		log.Fatalln(err)
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		log.Fatalln(err)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "combinePdfs")
	log.Printf("Saving... %v", "save.pdf")

	for _, fi := range fileInfos {
		if strings.HasSuffix(fi.Name(), ".pdf") {
			log.Printf("File name %v", fi.Name())
			_, r, err := pdfL.Open(fi.Name())
			if err != nil {
				log.Fatalln(err)
			}
			totalPage := r.NumPage()
			log.Printf("num of pages:  %v", totalPage)

			for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
				log.Printf("index %v", pageIndex)
				p := r.Page(pageIndex)
				if p.V.IsNull() {
					continue
				}
				pdf.AddPage()
				pdf.SetFont("Arial", "B", 16)
				txt, err := p.GetPlainText(nil)
				if err != nil {
					fmt.Printf("error: ", err)
				}
				pdf.Cell(40, 10, txt)
			}
		}
	}
	err = pdf.OutputFileAndClose("saved.pdf")
}
