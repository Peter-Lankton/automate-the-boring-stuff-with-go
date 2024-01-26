package main

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/ledongthuc/pdf"
)

func main() {

	content, err := readPdf("basic.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Garamond", "B", 16)
	pdf.Cell(40, 10, content)
	err = pdf.OutputFileAndClose("hello.pdf")

	return
}

func readPdf(path string) (string, error) {
	_, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	var textBuilder bytes.Buffer
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		s, err := p.GetPlainText(nil)
		if err != nil {
			fmt.Printf("Error: ", err)
		}

		textBuilder.WriteString(s)
	}
	return textBuilder.String(), nil
}
