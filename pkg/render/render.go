package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"github.com/ihorallin/bookings/pkg/config"
	"github.com/ihorallin/bookings/pkg/modules"
)

var functions = template.FuncMap {

}

var app *config.AppConfig

// NewTemplates sets the config form the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *modules.TemplateData) *modules.TemplateData {
	return td
}

// RenderTemplate render template using html/mplate
func RenderTemplate(w http.ResponseWriter, tmpl string, td *modules.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// Get the template cache from the App config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}


	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
