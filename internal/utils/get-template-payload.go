package utils

import "github.com/deriveddotdev/derived-go-template-parser/models"

// GetResolvePayload processes a template array and identifies duplicate paths
func GetResolvePayload(templateArray []models.TemplateInput) []models.Template {
	existing := make(map[int]bool)
	result := make([]models.Template, 0)

	for _, curr := range templateArray {
		if existing[curr.ID] {
			continue
		}

		// Find templates with duplicate paths
		duplicatePaths := findDuplicatePaths(templateArray, curr)

		// Mark all duplicate IDs as existing
		for _, dup := range duplicatePaths {
			existing[dup.ID] = true
		}

		if len(duplicatePaths) > 0 {
			// Add current template to duplicates list
			duplicatePaths = append(duplicatePaths, curr)

			// Create same path templates slice
			samePathTemplates := make([]models.TemplateList, 0, len(duplicatePaths))
			for _, dup := range duplicatePaths {
				samePathTemplates = append(samePathTemplates, models.TemplateList{
					Template: dup.Template,
					Pattern:  dup.Pattern,
				})
			}

			result = append(result, models.Template{
				ID:                curr.ID,
				Template:          curr.Template,
				Path:              curr.Path,
				TemplateType:      curr.TemplateType,
				Pattern:           curr.Pattern,
				SamePathTemplates: samePathTemplates,
				HasDuplicatePaths: true,
			})
		} else {
			result = append(result, models.Template{
				ID:                curr.ID,
				Template:          curr.Template,
				Path:              curr.Path,
				TemplateType:      curr.TemplateType,
				Pattern:           curr.Pattern,
				SamePathTemplates: []models.TemplateList{},
				HasDuplicatePaths: false,
			})
		}
	}

	return result
}

// findDuplicatePaths finds all templates with the same path and type as the current template (excluding current)
func findDuplicatePaths(templateArray []models.TemplateInput, curr models.TemplateInput) []models.TemplateInput {
	duplicates := make([]models.TemplateInput, 0)

	for _, item := range templateArray {
		if item.Path != "" &&
			item.Path == curr.Path &&
			item.TemplateType == curr.TemplateType &&
			item.ID != curr.ID {
			duplicates = append(duplicates, item)
		}
	}

	return duplicates
}

// PreparePayload creates the initial payload from template list
func PreparePayload(templateList []models.TemplateInput) []models.Template {
	payload := make([]models.Template, 0, len(templateList))

	for _, item := range templateList {
		payload = append(payload, models.Template{
			ID:                item.ID,
			Template:          item.Template,
			Path:              item.Path,
			TemplateType:      item.TemplateType,
			Pattern:           item.Pattern,
			SamePathTemplates: []models.TemplateList{},
			HasDuplicatePaths: false,
		})
	}
	return payload
}
