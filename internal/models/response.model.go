package models

type ResolveTemplatesResponse struct {
	Success bool               `json:"success"`
	Files   []ResolvedTemplate `json:"files"`
}

type GenerateFilesResponse struct {
	AddResults    []string `json:"addResults"`
	ModifyResults []string `json:"modifyResults"`
	Errors        []string `json:"errors"`
}

type ResolvePathResponse struct {
	Success bool           `json:"success"`
	Path    []ResolvedPath `json:"path"`
}

type ErrorResponse struct {
	Err        string `json:"err"`
	LineNumber string `json:"lineNumber"`
}

type SingleTemplateResolvedResponse struct {
	Success  bool          `json:"success"`
	Template string        `json:"template"`
	Error    ErrorResponse `json:"error"`
}

type GenerateTemplateWithAiResponse struct {
	Success  bool   `json:"success"`
	Template string `json:"template"`
	Error    string `json:"error"`
}
