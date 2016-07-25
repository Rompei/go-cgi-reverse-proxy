package main

import (
	"text/template"
)

// TemplateModel is model of template engine.
type TemplateModel struct {
	BaseURL string
	LogFile string
	Server  string
}

// NewTemplateModel is constructor of TemplateModel.
func NewTemplateModel(baseURL, server, logFile string) *TemplateModel {
	return &TemplateModel{
		BaseURL: baseURL,
		LogFile: logFile,
		Server:  server,
	}
}

func loadTemplateFromBinary(name string) (*template.Template, error) {
	tmplBin, err := Asset(name)
	if err != nil {
		return nil, err
	}
	return template.Must(template.New("t").Parse(string(tmplBin))), nil
}
