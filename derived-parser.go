package derivedgotemplateparser

import (
	"github.com/deriveddotdev/derived-go-template-parser/internal/models"
	"github.com/deriveddotdev/derived-go-template-parser/internal/resolvers"
)

type TemplateInput struct {
	ID           int    `json:"id"`
	Template     string `json:"template"`
	Path         string `json:"path"`
	TemplateType string `json:"templateType"`
	Pattern      string `json:"pattern"`
	FileContent  string `json:"fileContent"`
	Enabled      string `json:"enabled"`
}

func ResolveMultipleTemplates(templatesInput []models.TemplateInput, data interface{}) models.ResolveTemplatesResponse {
	return resolvers.ResolveMultipleTemplates(templatesInput, data)
}
