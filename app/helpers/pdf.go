package helpers

import (
	"errors"
	"log"
	"os"
	"path"
	"strings"
	"time"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

var (
	pdfPath = path.Join("public", "pdf")
)

type (
	// IPdfGenerator interface
	IPdfGenerator interface {
		HTMLToPdf(html string, filename string) error
		Find(filename string) (string, error)
	}
	// PdfGenerator struct
	PdfGenerator struct {
		generator *wkhtmltopdf.PDFGenerator
	}
)

// NewPdfGenerator pdf generator constructor
func NewPdfGenerator() IPdfGenerator {
	generator, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		log.Fatal(err)
	}

	gen := &PdfGenerator{generator}

	return gen
}

// HTMLToPdf writes html string to pdf and stores file on disk
func (gen *PdfGenerator) HTMLToPdf(html string, filename string) error {
	start := time.Now()
	if _, err := os.Stat(pdfPath); os.IsNotExist(err) {
		os.MkdirAll(pdfPath, 755)
	}

	pdfFilePath := path.Join(pdfPath, filename)

	gen.generator.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))
	err := gen.generator.Create()
	if err != nil {
		return err
	}

	err = gen.generator.WriteFile(pdfFilePath)
	if err != nil {
		return err
	}

	elapsed := time.Since(start)
	log.Printf("PDF generation took %s", elapsed)

	return nil
}

// Find finds pdf on disk and returns file path, otherwise returns error
func (gen *PdfGenerator) Find(filename string) (string, error) {
	pdfFilePath := path.Join(pdfPath, filename)

	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		return "", errors.New("file not found")
	}

	return pdfFilePath, nil
}
