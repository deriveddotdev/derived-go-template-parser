package core

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
	utils "github.com/deriveddotdev/derived-go-template-parser/internal/utils"
)

func Parser(tmplStr string) (*template.Template, error) {
	funcs := template.FuncMap{
		"titleCase":    utils.TitleCase,
		"camelCase":    utils.CamelCase,
		"snakeCase":    utils.SnakeCase,
		"kebabCase":    utils.KebabCase,
		"upperCase":    utils.UpperCase,
		"lowerCase":    utils.LowerCase,
		"constantCase": utils.ConstantCase,
		"pascalCase":   utils.PascalCase,
		"flatCase":     utils.FlatCase,
		"add":          func(a, b int) int { return a + b },
		"sub":          func(a, b int) int { return a - b },
		"mul":          func(a, b int) int { return a * b },
		"divide":       func(a, b int) int { return a / b },
		"isLast":       utils.IsLast,
		"sliceHas":     utils.SliceHas,
	}

	mergedFuncs := sprig.FuncMap()

	for key, fn := range funcs {
		mergedFuncs[key] = fn
	}

	return template.New("").Funcs(mergedFuncs).Parse(tmplStr)
}
