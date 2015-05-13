package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

var templates *template.Template

// custom template delimiters since the Go default delimiters clash
// with Angular's default.
var templateDelims = []string{"{{%", "%}}"}

func init() {
	// initialize the templates,
	// couldn't have used http://golang.org/pkg/html/template/#ParseGlob
	// since we have custom delimiters.
	basePath := "resources/templates/"
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// don't process folders themselves
		if info.IsDir() {
			return nil
		}
		templateName := path[len(basePath):]
		if templates == nil {
			templates = template.New(templateName)
			templates.Delims(templateDelims[0], templateDelims[1])
			_, err = templates.ParseFiles(path)
		} else {
			_, err = templates.New(templateName).ParseFiles(path)
		}
		log.Printf("Processed template %s\n", templateName)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}
