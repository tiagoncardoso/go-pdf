package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/tiagoncardoso/go-pdf/config"
	"github.com/tiagoncardoso/go-pdf/internal/application/usecase"
	"github.com/tiagoncardoso/go-pdf/internal/infra/http/types"
	"github.com/tiagoncardoso/go-pdf/pkg/logger"
)

type PdfHandler struct {
	ctx                  context.Context
	pdfGeneratorUsecase  *usecase.GeneratePdfFromHtml
	sendToStorageUsecase *usecase.SendFileToStorage
	generatePdfTempLink  *usecase.GenerateTempFileLink
	deleteTempFile       *usecase.DeleteTempFile
}

func NewPdfHandler(ctx context.Context, envConfig *config.EnvConfig) *PdfHandler {
	return &PdfHandler{
		ctx:                  ctx,
		pdfGeneratorUsecase:  usecase.NewGeneratePdfFromHtml(envConfig),
		sendToStorageUsecase: usecase.NewSendFileToStorage(envConfig),
		generatePdfTempLink:  usecase.NewGenerateTempFileLink(envConfig),
		deleteTempFile:       usecase.NewDeleteTempFile(envConfig),
	}
}

func (p *PdfHandler) GeneratePdf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	reportPath := chi.URLParam(r, "reportPath")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	htmlBody := r.FormValue("body")
	if htmlBody == "" {
		http.Error(w, "Missing 'body' form value", http.StatusBadRequest)
		return
	}

	htmlHeader := r.FormValue("header")
	if htmlHeader == "" {
		logger.Warn("Missing 'header' form value. Setting header to empty string.")
		htmlHeader = ""
	}

	pdfName, err := p.pdfGeneratorUsecase.Execute(htmlBody, htmlHeader)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate PDF: %v", err), http.StatusBadGateway)
		return
	}

	_, err = p.sendToStorageUsecase.Execute(reportPath, pdfName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send file to storage: %v", err), http.StatusBadGateway)
		return
	}

	err = p.deleteTempFile.Execute(pdfName)
	if err != nil {
		logger.Warn("Failed to delete temporary file.", "err", err)
	}

	err = json.NewEncoder(w).Encode(types.HttpOkResponse{
		Message: "PDF file generated and sent to storage successfully",
		Data: map[string]string{
			"fileID": strings.Replace(pdfName, ".pdf", "", 1),
		},
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

	reportPath := chi.URLParam(r, "reportPath")
	fileId := filepath.Join(reportPath, chi.URLParam(r, "fileId")+".pdf")

	link, err := p.generatePdfTempLink.Execute(fileId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate temporary link: %v", err), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(types.HttpOkResponse{
		Message: "Temporary link generated successfully ",
		Data: map[string]string{
			"link": link,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
