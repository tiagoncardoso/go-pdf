package usecase

import (
	"github.com/tiagoncardoso/go-pdf/config"
	"github.com/tiagoncardoso/go-pdf/internal/application/helpers"
	"github.com/tiagoncardoso/go-pdf/pkg/logger"
	pdfgen "github.com/tiagoncardoso/go-pdf/pkg/pdf-generator"
)

type GeneratePdfFromHtml struct {
	env *config.EnvConfig
}

func NewGeneratePdfFromHtml(env *config.EnvConfig) *GeneratePdfFromHtml {
	return &GeneratePdfFromHtml{
		env,
	}
}

func (p *GeneratePdfFromHtml) Execute(htmlContent string) (string, error) {
	pdfName := helpers.GenerateFileName(p.env.ReportPrefix)

	pdfGenerator := pdfgen.New(
		pdfgen.WithHTMLContent(htmlContent),
		pdfgen.WithOutputFilePath(p.env.OutputPath, pdfName),
		pdfgen.WithDPISet(p.env.Dpi),
		pdfgen.WithPageSizeSet(p.env.PageSize),
		pdfgen.WithOrientationSet(p.env.Orientation),
		pdfgen.WithTitle(p.env.Title),
	)

	pdfByte, err := pdfGenerator.GeneratePDF()
	if err != nil {
		logger.Error("error generating pdf file", "error", err.Error())
	}

	if err := pdfGenerator.CreateFile(pdfByte); err != nil {
		logger.Error("error creating pdf file", "error", err.Error())
	}

	return pdfName, nil
}
