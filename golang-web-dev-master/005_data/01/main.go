package main

import (
	"log"
	"os"
	"text/template"
)

type student struct {
	name  string
	email string
	age   int16
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	student := student{
		name:  "Andrew",
		email: "andrew@gmail.com",
		age:   18,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", student)
	if err != nil {
		log.Fatalln(err)
	}
}
