package view

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Info contains data to work with templates.
type Info struct {
	BaseURI string `yaml:"BaseURI"`

	Templates struct {
		Root      string   `yaml:"Root"`
		Children  []string `yaml:"Children"`
		Extension string   `yaml:"Extension"`
	} `yaml:"Templates"`

	Directories struct {
		Assets string `yaml:"Assets"`
		Views  string `yaml:"Views"`
	} `yaml:"Directories"`

	Vars       map[string]interface{}
	extensions template.FuncMap
	modifiers  []ModifyFunc

	templates []string
}

// TODO: figure out concurrency.

// New accepts templates and returnse a view.
func (v *Info) New(templateList ...string) *Info {
	v.Vars = make(map[string]interface{})
	v.templates = append(v.templates, templateList...)

	return v
}

// Render parses templates.
func (v *Info) Render(w http.ResponseWriter, r *http.Request) error {
	v.templates = append([]string{v.Templates.Root}, v.templates...)
	v.templates = append(v.templates, v.Templates.Children...)

	for i, tmpl := range v.templates {
		path, _ := filepath.Abs(v.Directories.Views + string(os.PathSeparator) + tmpl + v.Templates.Extension)
		v.templates[i] = path
	}

	extensions := v.extensions
	templates, err := template.New("changeme").Funcs(extensions).ParseFiles(v.templates...)
	if err != nil {
		log.Println(errors.New("template parse error: " + err.Error()))
		http.Error(w, "TODO: write funny error for user", http.StatusInternalServerError)
		return err
	}

	for _, fn := range v.modifiers {
		fn(w, r, v)
	}

	if err = templates.Funcs(extensions).ExecuteTemplate(w, v.Templates.Root+v.Templates.Extension, v.Vars); err != nil {
		log.Println(errors.New("template error: " + err.Error()))
		http.Error(w, "TODO: write funny error for user", http.StatusInternalServerError)
	}

	return err
}

// ModifyFunc modify the view before rendering.
type ModifyFunc func(http.ResponseWriter, *http.Request, *Info)

// Modifiers set the modifiers for the View that run before rendering.
func (v *Info) Modifiers(fn ...ModifyFunc) {
	v.modifiers = fn
}

// Extensions combines all template.FuncMaps into one map.
func (v *Info) Extensions(fms ...template.FuncMap) {
	fm := make(template.FuncMap)

	for _, m := range fms {
		for k, val := range m {
			fm[k] = val
		}
	}

	v.extensions = fm
}
