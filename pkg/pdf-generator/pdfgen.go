package pdf_generator

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PdfParams struct {
	Orientation string
	PageSize    string
}

type PDFGenerator struct {
	htmlContent string
}

func NewPDFGenerator(htmlContent string) *PDFGenerator {
	return &PDFGenerator{
		htmlContent: htmlContent,
	}
}

func (p *PDFGenerator) Generate(pdfParams PdfParams) ([]byte, error) {
	pdfGen, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfContent := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(p.htmlContent)))
	pdfGen.AddPage(pdfContent)

	err = pdfGen.Create()
	if err != nil {
		return nil, err
	}

	p.SetParams(pdfGen, pdfParams)

	return pdfGen.Bytes(), nil
}

func (p *PDFGenerator) SetParams(pdf *wkhtmltopdf.PDFGenerator, params PdfParams) {
	p.SetPageSize(params.PageSize)
	pdf.Orientation.Set(p.SetPageSize(params.PageSize))
	pdf.Orientation.Set(p.SetOrientation(params.Orientation))

	pdf.Dpi.Set(300)
	pdf.Title.Set("Generated PDF")
}

func (p *PDFGenerator) SetPageSize(pageSize string) string {
	if pageSize != "" {
		return pageSize
	}

	return wkhtmltopdf.PageSizeA4
}

func (p *PDFGenerator) SetOrientation(orientation string) string {
	if orientation != "" {
		return orientation
	}

	return wkhtmltopdf.OrientationPortrait
}
