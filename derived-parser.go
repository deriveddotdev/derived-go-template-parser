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

func ResolveMultipleTemplates(templatesInput []TemplateInput, data interface{}) models.ResolveTemplatesResponse {
	internalInputs := make([]models.TemplateInput, len(templatesInput))
	for i, t := range templatesInput {
		internalInputs[i] = models.TemplateInput{
			ID:           t.ID,
			Template:     t.Template,
			Path:         t.Path,
			TemplateType: t.TemplateType,
			Pattern:      t.Pattern,
			FileContent:  t.FileContent,
			Enabled:      t.Enabled,
		}
	}
	return resolvers.ResolveMultipleTemplates(internalInputs, data)
}
