package core

import (
	"bytes"
)

func ResolveTemplateString(tmplStr string, data interface{}) (string, error) {
	tmpl, err := Parser(tmplStr)
	if err != nil {
		return "", err
	}

	var output bytes.Buffer
	err = tmpl.Execute(&output, data)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
