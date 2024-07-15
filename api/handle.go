package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

var (
	credentialFile = "./api-go-429422-28699526298f.json"
)

func HandleExtractRG(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL não encontrada", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	credential, err := os.ReadFile(credentialFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao ler o arquivo: %v", err), http.StatusInternalServerError)
		return
	}

	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsJSON(credential))
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao criar cliente da Vision API: %v", err), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao abrir URL do RG: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	image := &visionpb.Image{
		Source: &visionpb.ImageSource{
			ImageUri: url,
		},
	}

	annotations, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao detectar texto na imagem: %v", err), http.StatusInternalServerError)
		return
	}

	var extractedText []string
	for _, page := range annotations.Pages {
		for _, block := range page.Blocks {
			for _, paragraph := range block.Paragraphs {
				var paragraphText string
				for _, word := range paragraph.Words {
					var wordText string
					for _, symbol := range word.Symbols {
						wordText += symbol.Text
					}
					paragraphText += wordText + " "
				}
				extractedText = append(extractedText, paragraphText)
			}
		}
	}

	responseJSON, err := json.Marshal(extractedText)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao converter o JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func HandleExtractCNH(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL não encontrada", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	credential, err := os.ReadFile(credentialFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao ler arquivo: %v", err), http.StatusInternalServerError)
		return
	}

	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsJSON(credential))
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao criar cliente da Vision API: %v", err), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao abrir URL da CNH: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	image := &visionpb.Image{
		Source: &visionpb.ImageSource{
			ImageUri: url,
		},
	}

	annotations, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao detectar texto da imagem: %v", err), http.StatusInternalServerError)
		return
	}

	var extractedText []string
	for _, page := range annotations.Pages {
		for _, block := range page.Blocks {
			for _, paragraph := range block.Paragraphs {
				var paragraphText string
				for _, word := range paragraph.Words {
					var wordText string
					for _, symbol := range word.Symbols {
						wordText += symbol.Text
					}
					paragraphText += wordText + " "
				}
				extractedText = append(extractedText, paragraphText)
			}
		}
	}

	responseJSON, err := json.Marshal(extractedText)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao converter o JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func HandleExtractPassport(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL não encontrada", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	credentital, err := os.ReadFile(credentialFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao ler arquivo: %v", err), http.StatusInternalServerError)
		return
	}

	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsJSON(credentital))
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao criar cliente da Vision API: %v", err), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao abrir imagem da URL do Passaporte: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	image := &visionpb.Image{
		Source: &visionpb.ImageSource{
			ImageUri: url,
		},
	}

	annotations, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao detectar texto da imagem: %v", err), http.StatusInternalServerError)
		return
	}

	var extractedText []string
	for _, page := range annotations.Pages {
		for _, block := range page.Blocks {
			for _, paragraph := range block.Paragraphs {
				var paragraphText string
				for _, word := range paragraph.Words {
					var wordText string
					for _, symbol := range word.Symbols {
						wordText += symbol.Text
					}
					paragraphText += wordText + " "
				}
				extractedText = append(extractedText, paragraphText)
			}
		}
	}

	responseJSON, err := json.Marshal(extractedText)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao converter JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
