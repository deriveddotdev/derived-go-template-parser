package resolvers

import (
	"regexp"

	"github.com/deriveddotdev/derived-go-template-parser/internal/core"
	"github.com/deriveddotdev/derived-go-template-parser/internal/models"
)

func ResolveTemplate(template string, data interface{}) models.SingleTemplateResolvedResponse {
	resolvedTemplate, err := core.ResolveTemplateString(template, data)
	if err != nil {
		re := regexp.MustCompile(`template: :(\d+):`)
		matches := re.FindStringSubmatch(err.Error())
		if len(matches) > 1 {
			return models.SingleTemplateResolvedResponse{
				Success: false,
				Error: models.ErrorResponse{
					Err:        err.Error(),
					LineNumber: matches[1],
				},
			}
		}
		return models.SingleTemplateResolvedResponse{
			Success: false,
			Error: models.ErrorResponse{
				Err:        err.Error(),
				LineNumber: "0",
			},
		}
	}
	return models.SingleTemplateResolvedResponse{
		Success:  true,
		Template: resolvedTemplate,
	}
}
