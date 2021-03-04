// Package link provides a funcmap for html/template to generate a hyperlink.
package link

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Map returns a template.FuncMap that returns hyperlink tags.
func Map(baseURI string) template.FuncMap {
	f := make(template.FuncMap)

	f["JS"] = func(fpath string) template.HTML {
		path, err := assetTimePath(baseURI, fpath)

		if err != nil {
			log.Println("JS Error:", err)
			return template.HTML("<!-- JS Error: " + fpath + " -->")
		}

		return template.HTML(`<script type="text/javascript" src="` + path + `"></script>`)
	}

	f["CSS"] = func(fpath, media string) template.HTML {
		path, err := assetTimePath(baseURI, fpath)

		if err != nil {
			log.Println("CSS Error:", err)
			return template.HTML("<!-- CSS Error: " + fpath + " -->")
		}

		return template.HTML(fmt.Sprintf(`<link media="%v" rel="stylesheet" type="text/css" href="%v" />`, media, path))
	}

	return f
}

// assetTimePath returns a URL with the proper base URI and timestamp appended.
// Works for CSS and JS assets and determines if local or on the web by the
// number of slashes at the beginning of the string. A prefix of // is web and
// / is local.
func assetTimePath(baseURI, resource string) (string, error) {
	if strings.HasPrefix(resource, "//") {
		return resource, nil
	}

	resource = strings.TrimLeft(resource, "/")

	abs, err := filepath.Abs(filepath.Join("", resource))
	if err != nil {
		return "", err
	}

	time, err := fileTime(abs)
	if err != nil {
		return "", err
	}

	return baseURI + resource + "?" + time, nil
}

// fileTime returns the modification time of the file.
func fileTime(name string) (string, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return "", err
	}

	mtime := fi.ModTime().Unix()

	return fmt.Sprintf("%v", mtime), nil
}
