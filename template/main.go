package main

import (
	"os"
	"text/template"
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
			Name:   "Jujube",
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

	var tmplFile = "lastpet.tmpl"
	tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err  = tmpl.Execute(os.Stdout,dogs)
	if err != nil {
		panic(err)
	}
} 
