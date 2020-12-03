package main

import (
	"fmt"
	"log"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	url := "http://192.168.2.12/pandabi/#/release/e2c9816cd002f7d30afff7cf9a8e0aba"

	pdfg.AddPage(wkhtmltopdf.NewPage(url))

	// PDFط³
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// ³
	err = pdfg.WriteFile("./google.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tada!")
}
