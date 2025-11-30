package resolvers

import (
	"github.com/deriveddotdev/derived-go-template-parser/internal/core"
	"github.com/deriveddotdev/derived-go-template-parser/models"
)

func ResolvePathList(path []models.Path, data interface{}) models.ResolvePathResponse {
	resolvedPaths := []models.ResolvedPath{}
	for _, p := range path {
		resolvedPath, err := core.ResolveTemplateString(p.Path, data)
		if err != nil {
			return models.ResolvePathResponse{
				Success: false,
				Path:    []models.ResolvedPath{},
			}
		}
		resolvedPaths = append(resolvedPaths, models.ResolvedPath{
			TemplateID:   p.TemplateID,
			Path:         p.Path,
			ResolvedPath: resolvedPath,
		})
	}
	return models.ResolvePathResponse{
		Success: true,
		Path:    resolvedPaths,
	}
}
