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
	htmlHeader     string
	OutputPath     string
	fullOutputPath string
	Dpi            uint
	PageSize       string
	Title          string
	Orientation    string
	Pagination     bool
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

	if p.params.Pagination {
		pdfContent.FooterLine.Set(true)
		pdfContent.FooterRight.Set("Página [page] de [toPage]")
		pdfContent.FooterFontSize.Set(9)
		pdfContent.FooterSpacing.Set(5)
	}

	var headerFile *os.File
	headerFile, err = os.CreateTemp("./", "header-*.html")
	if err != nil {
		return nil, err
	}
	defer os.Remove(headerFile.Name())

	if p.params.htmlHeader != "" {

		if _, err := headerFile.Write([]byte(p.params.htmlHeader)); err != nil {
			return nil, err
		}
		headerFile.Close()
		pdfContent.HeaderHTML.Set(headerFile.Name())
		pdfContent.HeaderSpacing.Set(5)
	}

	pdfGen.AddPage(pdfContent)

	pdfGen.Dpi.Set(p.params.Dpi)
	pdfGen.MarginTop.Set(25)    // Set top margin to 20mm for header space
	pdfGen.MarginBottom.Set(15) // Set bottom margin as needed
	pdfGen.PageSize.Set(p.params.PageSize)
	pdfGen.Orientation.Set(p.params.Orientation)
	pdfGen.Title.Set(p.params.Title)

	err = pdfGen.Create()
	if err != nil {
		return nil, err
	}

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

func WithHeaderHTMLContent(htmlHeader string) Option {
	return func(p *PDFGenerator) {
		p.params.htmlHeader = htmlHeader
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
			p.params.fullOutputPath = filepath.Join("../internal/output/", strconv.FormatInt(now, 10)+".pdf")
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

func WithPaginationSet(pagination bool) Option {
	return func(p *PDFGenerator) {
		p.params.Pagination = pagination
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
