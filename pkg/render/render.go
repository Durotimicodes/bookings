package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Durotimicodes/bookings/pkg/config"
	"github.com/Durotimicodes/bookings/pkg/model"
)

var app *config.AppConfig

//NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *model.TemplateData) *model.TemplateData {
	return td
}

//Render template using html template
func RenderTemplate(w http.ResponseWriter, html string, td *model.TemplateData) {

	//get the template cache from the app config

	//ADVANCED METHOD TO CACHE FILES
	//a. create a template cache

	var tc map[string]*template.Template
	if app.UseCache { // if use cache is true read the template from the template cache
		tc = app.TemplateCache
	} else { //otherwise create a new template/rebuild the template cache
		tc, _ = CreateTemplateCache()
	}

	//b. get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	//create a buffer the holds bytes
	buf := new(bytes.Buffer)

	//add default template
	td = AddDefaultData(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//c. render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	//data structure to store values
	myCache := map[string]*template.Template{}

	//get all of the files names with .page.html from ./templates
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}
	//range through all files that end with *.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil

}
