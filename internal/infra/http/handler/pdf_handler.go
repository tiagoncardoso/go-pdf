package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tiagoncardoso/go-pdf/config"
	"github.com/tiagoncardoso/go-pdf/internal/application/usecase"
	"github.com/tiagoncardoso/go-pdf/internal/infra/http/types"
)

type PdfHandler struct {
	ctx                  context.Context
	pdfGeneratorUsecase  *usecase.GeneratePdfFromHtml
	sendToStorageUsecase *usecase.SendFileToStorage
}

func NewPdfHandler(ctx context.Context, envConfig *config.EnvConfig) *PdfHandler {
	return &PdfHandler{
		ctx:                  ctx,
		pdfGeneratorUsecase:  usecase.NewGeneratePdfFromHtml(envConfig),
		sendToStorageUsecase: usecase.NewSendFileToStorage(envConfig),
	}
}

func (p *PdfHandler) GeneratePdf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = p.pdfGeneratorUsecase.Execute(string(body))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate PDF: %v", err), http.StatusInternalServerError)
		return
	}
	// TODO: Chamar Caso de Uso com Parametros do PDF + Body para gerar o PDF
	// TODO: Utilizar o PDF gerado na resposta da requisição

	err = p.sendToStorageUsecase.Execute("caminho/do/arquivo.pdf")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send file to storage: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(types.HttpOkResponse{
		Message: "PDF generation not implemented yet",
		Data:    nil,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
