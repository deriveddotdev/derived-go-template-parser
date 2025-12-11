package resolvers

import (
	"crypto/sha1"
	"log"
	"regexp"
	"strings"

	"github.com/deriveddotdev/derived-go-template-parser/internal/core"
	"github.com/deriveddotdev/derived-go-template-parser/internal/models"
	"github.com/deriveddotdev/derived-go-template-parser/internal/utils"
)

type ModifyResult struct {
	Content     string
	FileContent string
}

func ResolveMultipleTemplates(templatesInput []models.TemplateInput, data interface{}) models.ResolveTemplatesResponse {
	resolvedTemplates := []models.ResolvedTemplate{}

	templates := utils.GetResolvePayload(templatesInput)

	for _, template := range templates {
		if template.HasDuplicatePaths {
			resolvedTemplates = append(resolvedTemplates, handleDuplicatePaths(template, data))
		} else {
			resolvedTemplates = append(resolvedTemplates, handleDistinctPaths(template, data))
		}
	}
	return models.ResolveTemplatesResponse{
		Success: true,
		Files:   resolvedTemplates,
	}
}

func handleDistinctPaths(template models.Template, data interface{}) models.ResolvedTemplate {
	resolvedTemplate, err := core.ResolveTemplateString(template.Template, data)
	resolvedPath, pathErr := core.ResolveTemplateString(template.Path, data)
	resolvedEnabled, enabledErr := core.ResolveTemplateString(template.Enabled, data)
	result := models.ResolvedTemplate{
		ID:                  template.ID,
		Template:            template.Template,
		ResolvedTemplate:    resolvedTemplate,
		ResolvedPath:        resolvedPath,
		Path:                template.Path,
		TemplateError:       "",
		PathError:           "",
		FileContent:         template.FileContent,
		ModifiedFileContent: template.FileContent,
		TemplateType:        template.TemplateType,
		LineNumbers:         []int{},
		IsNotEnabled:        resolvedEnabled == "false",
	}
	if pathErr != nil {
		result.PathError = pathErr.Error()
	}

	if err != nil {
		result.TemplateError = err.Error()
	}

	if enabledErr != nil {
		result.IsNotEnabled = false
	}

	if template.TemplateType == models.TemplateTypeModify {
		modifiedResult := modifyFile(template.Pattern, result.FileContent, resolvedTemplate)
		result.ModifiedFileContent = modifiedResult.Content
		result.LineNumbers = getAddedLines(result.FileContent, result.ModifiedFileContent)
	}
	return result
}

func handleDuplicatePaths(template models.Template, data interface{}) models.ResolvedTemplate {
	log.Println("file content", template.FileContent)
	resolvedPath, pathErr := core.ResolveTemplateString(template.Path, data)
	result := models.ResolvedTemplate{
		ID:                  template.ID,
		Template:            template.Template,
		ResolvedTemplate:    "",
		ResolvedPath:        resolvedPath,
		Path:                template.Path,
		TemplateError:       "",
		PathError:           "",
		FileContent:         template.FileContent,
		ModifiedFileContent: template.FileContent,
		TemplateType:        template.TemplateType,
		LineNumbers:         []int{},
	}
	if pathErr != nil {
		result.PathError = pathErr.Error()
	}

	for _, samePathTemplate := range template.SamePathTemplates {
		resolvedTemplate, _ := core.ResolveTemplateString(samePathTemplate.Template, data)
		modifiedResult := modifyFile(samePathTemplate.Pattern, result.ModifiedFileContent, resolvedTemplate)
		result.ModifiedFileContent = modifiedResult.Content
	}

	result.LineNumbers = getAddedLines(result.FileContent, result.ModifiedFileContent)
	return result
}

func modifyFile(pattern string, fileContent string, resolvedTemplate string) ModifyResult {
	result := ModifyResult{
		Content: fileContent,
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return result
	}

	// Apply the replacement
	result.Content = re.ReplaceAllString(fileContent, resolvedTemplate)

	return result
}

func hashLine(line string) [sha1.Size]byte {
	return sha1.Sum([]byte(line))
}

func getAddedLines(oldContent, newContent string) []int {
	oldLines := make(map[[sha1.Size]byte]struct{})

	// Store hashed lines from old content in map
	for _, line := range strings.Split(oldContent, "\n") {
		hash := hashLine(line)
		oldLines[hash] = struct{}{}
	}

	// Find added lines in new content
	var addedLines []int
	newLines := strings.Split(newContent, "\n")
	for i, line := range newLines {
		hash := hashLine(line)
		if _, exists := oldLines[hash]; !exists {
			addedLines = append(addedLines, i+1)
		}
	}

	return addedLines
}
