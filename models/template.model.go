package models

const (
	TemplateTypeAdd    = "add"
	TemplateTypeModify = "modify"
)

type TemplateList struct {
	Pattern  string `json:"pattern"`
	Template string `json:"template"`
}

type TemplateInput struct {
	ID           int    `json:"id"`
	Template     string `json:"template"`
	Path         string `json:"path"`
	TemplateType string `json:"templateType"`
	Pattern      string `json:"pattern"`
	FileContent  string `json:"fileContent"`
	Enabled      string `json:"enabled"`
}

type Template struct {
	ID                int            `json:"id"`
	Template          string         `json:"template"`
	Path              string         `json:"path"`
	TemplateType      string         `json:"templateType"`
	Pattern           string         `json:"pattern"`
	SamePathTemplates []TemplateList `json:"samePathTemplates"`
	HasDuplicatePaths bool           `json:"hasDuplicatePaths"`
	FileContent       string         `json:"fileContent"`
	Enabled           string         `json:"enabled"`
}

type ResolvedTemplate struct {
	ID                  int    `json:"id"`
	Template            string `json:"template"`
	Path                string `json:"path"`
	TemplateType        string `json:"templateType"`
	ResolvedTemplate    string `json:"resolvedTemplate"`
	ResolvedPath        string `json:"resolvedPath"`
	FileContent         string `json:"fileContent"`
	ModifiedFileContent string `json:"modifiedFileContent"`
	PathError           string `json:"pathError"`
	TemplateError       string `json:"templateError"`
	LineNumbers         []int  `json:"lineNumbers"`
	IsNotEnabled        bool   `json:"isNotEnabled"`
}

type Path struct {
	TemplateID int    `json:"templateID"`
	Path       string `json:"path"`
}

type ResolvedPath struct {
	TemplateID   int    `json:"templateID"`
	Path         string `json:"path"`
	ResolvedPath string `json:"resolvedPath"`
	Error        string `json:"error"`
}

type ResolveTemplateStringPayload struct {
	Template string                 `json:"template"`
	Data     map[string]interface{} `json:"data"`
}
