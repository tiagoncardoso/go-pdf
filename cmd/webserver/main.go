package main

import (
	"context"

	"github.com/tiagoncardoso/go-pdf/config"
	"github.com/tiagoncardoso/go-pdf/internal/infra/http"
	"github.com/tiagoncardoso/go-pdf/internal/infra/http/handler"
)

func main() {
	ctx := context.Background()
	envConf, err := config.SetupEnvConfig()
	if err != nil {
		panic(err)
	}

	pdfGenHandler := handler.NewPdfHandler(ctx, envConf)

	webServer := http.NewWebServer(envConf.AppPort)

	webServer.AddHandler("/pdf/generate", "POST", pdfGenHandler.GeneratePdf)

	webServer.Start()
}
