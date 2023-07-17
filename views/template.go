package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
)

type Template struct {
	htmlTpl *template.Template
}

func ParseFS(fs fs.FS, patterns ...string) (*Template, error) {
	t, err := template.ParseFS(fs, patterns...)
	if err != nil {
		return &Template{}, fmt.Errorf("ParseFS %w", err)
	} else {
		return &Template{
			htmlTpl: t,
		}, nil
	}
}

func ParseTemplate(filename string) (*Template, error) {
	tplPath := filepath.Join("templates", filename)
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return &Template{}, fmt.Errorf("%w", err)
	} else {
		return &Template{
			htmlTpl: tpl,
		}, nil
	}
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		http.Error(w, "passed wrong data to template", http.
			StatusInternalServerError)
		panic(err)
	}
}

func Must(t *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return t
}
