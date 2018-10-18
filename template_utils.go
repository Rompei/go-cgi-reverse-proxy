package main

import (
	"text/template"

	"github.com/gobuffalo/packr"
)

// TemplateModel is model of template engine.
type TemplateModel struct {
	BaseURL string
	LogFile string
	Server  string
}

// NewTemplateModel is constructor of TemplateModel.
func NewTemplateModel(baseURL, server string) *TemplateModel {
	return &TemplateModel{
		BaseURL: baseURL,
		Server:  server,
	}
}

func loadTemplateFromBinary(name string) (*template.Template, error) {
	box := packr.NewBox("./templates")
	tmplBin, err := box.MustString(name)
	if err != nil {
		return nil, err
	}
	return template.Must(template.New("t").Parse(tmplBin)), nil
}
