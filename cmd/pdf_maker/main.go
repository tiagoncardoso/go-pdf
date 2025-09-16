package main

import (
	"github.com/tiagoncardoso/go-pdf/pkg/logger"
	"github.com/tiagoncardoso/go-pdf/pkg/pdf-generator"
)

func main() {
	pdfConfig := pdfgen.PdfParams{
		Dpi:         300,
		OutputPath:  "internal/output",
		Title:       "PDF Gerado",
		Orientation: "Portrait",
		PageSize:    "A4",
	}

	html := `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Test PDF</title>
<style>
    body {
      font-family: "Arial", sans-serif;
      font-size: 14pt;
    }
  </style>
</head>
<body>
  <h1 style="color:#000000">Hello, PDF!</h1>
</body>
</html>`

	pdfGenerator := pdfgen.New(
		pdfgen.WithHTMLContent(html),
		pdfgen.WithOutputFilePath(pdfConfig.OutputPath, "meupdf.pdf"),
		pdfgen.WithDPISet(pdfConfig.Dpi),
		pdfgen.WithPageSizeSet(pdfConfig.PageSize),
		pdfgen.WithOrientationSet(pdfConfig.Orientation),
		pdfgen.WithTitle(pdfConfig.Title),
	)

	pdfByte, err := pdfGenerator.GeneratePDF()
	if err != nil {
		logger.Error("error generating pdf file", "error", err.Error())
	}

	if err := pdfGenerator.CreateFile(pdfByte); err != nil {
		logger.Error("error creating pdf file", "error", err.Error())
	}
}
