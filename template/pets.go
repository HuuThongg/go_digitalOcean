package main

import (
	"os"
	"html/template"
	"strings"
)
type  Pet struct {
	Name string `json:"name"`
	Sex string `json:"sex"`
	Intact bool `json:"intact"`
	Age string `json:"age"`
	Breed []string `json:"breed"`
}

func main() {
	dogs := []Pet{
		{
			Name:   "<script>alert(\"Gotcha!\");</script>Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  []string{"German Shepherd", "Pit Bull"},
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "13 years, 3 months",
			Breed:  []string{"German Shepherd", "Border Collie"},
		},
		{
			Name:	"Bruce Wayne",
			Sex:	"Male",
			Intact:	false,
			Age:	"3 years, 8 months",
			Breed:  []string{"Chihuahua"},
		},
	}

	funcMap := template.FuncMap{
		"dec": func(i int) int { return i - 1},
		"replace": strings.ReplaceAll,
		"join":    strings.Join,
	}

	var tmplFile = "petsHtml.tmpl"
	tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	var f *os.File
	f, err  = os.Create("pets.html")
	if err != nil {
		panic(err)
	}

	err  = tmpl.Execute(f,dogs)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
} 
