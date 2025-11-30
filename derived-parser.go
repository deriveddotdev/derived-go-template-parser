package derivedgotemplateparser

import (
	"github.com/deriveddotdev/derived-go-template-parser/internal/models"
	"github.com/deriveddotdev/derived-go-template-parser/internal/resolvers"
)

func ResolveMultipleTemplates(templatesInput []models.TemplateInput, data interface{}) models.ResolveTemplatesResponse {
	return resolvers.ResolveMultipleTemplates(templatesInput, data)
}
