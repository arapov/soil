package view

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

// Info contains data to work with templates.
type Info struct {
	Template struct {
		Root      string   `yaml:"Root"`
		Children  []string `yaml:"Children"`
		Extension string   `yaml:"Extension"`
	} `yaml:"Template"`

	Directories struct {
		Assets string `yaml:"Assets"`
		Views  string `yaml:"Views"`
	} `yaml:"Directories"`

	templates []string
}

// New accepts templates and returnse a view.
func (v *Info) New(templateList ...string) *Info {
	v.templates = append(v.templates, templateList...)

	return v
}

// Render parses templates.
func (v *Info) Render(w http.ResponseWriter) error {
	v.templates = append([]string{v.Template.Root}, v.templates...)
	v.templates = append(v.templates, v.Template.Children...)

	for i, tmpl := range v.templates {
		path, _ := filepath.Abs(v.Directories.Views + string(os.PathSeparator) + tmpl + v.Template.Extension)
		v.templates[i] = path
	}

	templates, err := template.ParseFiles(v.templates...)
	if err != nil {
		log.Println(errors.New("template parse error: " + err.Error()))
		http.Error(w, "TODO: write funny error for user", http.StatusInternalServerError)
		return err
	}

	if err = templates.ExecuteTemplate(w, v.Template.Root+v.Template.Extension, nil); err != nil {
		log.Println(errors.New("template error: " + err.Error()))
		http.Error(w, "TODO: write funny error for user", http.StatusInternalServerError)
	}

	return err
}
