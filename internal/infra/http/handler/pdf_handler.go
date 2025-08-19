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
	generatePdfTempLink  *usecase.GenerateTempFileLink
}

func NewPdfHandler(ctx context.Context, envConfig *config.EnvConfig) *PdfHandler {
	return &PdfHandler{
		ctx:                  ctx,
		pdfGeneratorUsecase:  usecase.NewGeneratePdfFromHtml(envConfig),
		sendToStorageUsecase: usecase.NewSendFileToStorage(envConfig),
		generatePdfTempLink:  usecase.NewGenerateTempFileLink(envConfig),
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

	pdfName, err := p.pdfGeneratorUsecase.Execute(string(body))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate PDF: %v", err), http.StatusInternalServerError)
		return
	}

	objectKey, err := p.sendToStorageUsecase.Execute(pdfName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send file to storage: %v", err), http.StatusInternalServerError)
		return
	}

	// TODO: Melhorar resposta da API
	// TODO: Adicionar referência para que o arquivo possa ser baixado (via link temporário)
	// TODO: Gerar Hash do nome do arquivo para buscar no storage
	err = json.NewEncoder(w).Encode(types.HttpOkResponse{
		Message: "File " + objectKey + " generated and sent to storage successfully",
		Data:    nil,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (p *PdfHandler) GenerateTempLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//TODO: Gerar hash do nome do arquivo e buscar no storage
	// Obter a hash do nome do arquivo via query param
	link, err := p.generatePdfTempLink.Execute("analytics/sca-report_1755566931.pdf")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate temporary link: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(types.HttpOkResponse{
		Message: "Temporary link: " + link,
		Data:    nil,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
