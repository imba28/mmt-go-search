package api

import (
	"html/template"
	"io/ioutil"
	"log"
)

func initTemplate(template *template.Template, funcMap template.FuncMap) *template.Template {
	template.Funcs(funcMap)
	contents, err := ioutil.ReadFile("template/" + template.Name() + ".gohtml")
	if err != nil {
		log.Panic(err)
	}
	template.Parse(string(contents))
	return template
}
