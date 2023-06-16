package views

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type Template struct {
	htmlTpl *template.Template
}

func ParseTemplate(filename string) (Template, error) {
	tplPath := filepath.Join("templates", filename)
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	} else {
		return Template{
			htmlTpl: tpl,
		}, nil
	}
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		http.Error(w, "passed wrong data to template", http.
			StatusInternalServerError)
		return
	}
}
