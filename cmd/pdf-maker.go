package main

import (
	pdf_generator "github.com/tiagoncardoso/go-pdf/pkg/pdf-generator"
	"log"
	"os"
	"path/filepath"
)

func main() {
	pdfConfig := pdf_generator.PdfParams{
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

	pdfBytes, err := pdf_generator.NewPDFGenerator(html).Generate(pdfConfig)
	if err != nil {
		log.Fatal(err)
	}

	outputDir := "internal/output"
	outputPath := filepath.Join(outputDir, "output.pdf")

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(err)
	}

	if err := os.WriteFile(outputPath, pdfBytes, 0755); err != nil {
		panic(err)
	}
}
