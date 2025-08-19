package pdfgen

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PDFGenerator struct {
	params PdfParams
}

type PdfParams struct {
	htmlContent    string
	OutputPath     string
	fullOutputPath string
	Dpi            uint
	PageSize       string
	Title          string
	Orientation    string
}

type Option func(*PDFGenerator)

func New(options ...func(*PDFGenerator)) *PDFGenerator {
	pdfGen := &PDFGenerator{}

	for _, option := range options {
		option(pdfGen)
	}

	return pdfGen
}

func (p *PDFGenerator) GeneratePDF() ([]byte, error) {
	pdfGen, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfContent := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(p.params.htmlContent)))
	pdfGen.AddPage(pdfContent)

	err = pdfGen.Create()
	if err != nil {
		return nil, err
	}

	pdfGen.Dpi.Set(p.params.Dpi)
	pdfGen.PageSize.Set(p.params.PageSize)
	pdfGen.Orientation.Set(p.params.Orientation)
	pdfGen.Title.Set(p.params.Title)

	return pdfGen.Bytes(), nil
}

func (p *PDFGenerator) CreateFile(fileBytes []byte) error {
	if err := os.MkdirAll(p.params.OutputPath, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(p.params.fullOutputPath, fileBytes, 0755); err != nil {
		return err
	}

	return nil
}

func WithHTMLContent(htmlContent string) Option {
	return func(p *PDFGenerator) {
		p.params.htmlContent = htmlContent
	}
}

func WithOutputFilePath(filePath, filename string) Option {
	now := time.Now().Unix()
	return func(p *PDFGenerator) {
		if filePath != "" && filename != "" {
			p.params.OutputPath = filePath
			p.params.fullOutputPath = filepath.Join(filePath, filename)
		} else {
			p.params.OutputPath = filePath
			p.params.fullOutputPath = filepath.Join("../internal/output/output", strconv.FormatInt(now, 10)+".pdf")
		}
	}
}

func WithDPISet(dpi uint) Option {
	return func(p *PDFGenerator) {
		if dpi > 0 {
			p.params.Dpi = dpi
		} else {
			p.params.Dpi = 300
		}
	}
}

func WithPageSizeSet(pageSize string) Option {
	return func(p *PDFGenerator) {
		if pageSize != "" {
			p.params.PageSize = pageSize
		} else {
			p.params.PageSize = wkhtmltopdf.PageSizeA4
		}
	}
}

func WithOrientationSet(orientation string) Option {
	return func(p *PDFGenerator) {
		if orientation != "" {
			p.params.Orientation = orientation
		} else {
			p.params.Orientation = wkhtmltopdf.OrientationPortrait
		}
	}
}

func WithTitle(title string) Option {
	return func(p *PDFGenerator) {
		if title != "" {
			p.params.Title = title
		} else {
			p.params.Title = "Generated PDF"
		}
	}
}
