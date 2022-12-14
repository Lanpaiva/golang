package main

import (
	"html/template"
	"os"
)

type Estudo struct {
	Nome         string
	CargaHoraria int
}

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	// A função t já possui um tratamento de erro, template.New conseguimos fazer a busca do arquivo
	//Parse.Files faz leitura do arquivo com o os.Stdout setando no terminal.
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	// // templa := template.New("CursoGotemplate")
	// // templa, err := templa.Parse("Estudo, {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	// // if err != nil {
	// // 	panic(err)
	// // }

	err := t.Execute(os.Stdout, Cursos{
		{"Go:", 40},
		{"Node:", 70},
		{"Typescript:", 50},
	})
	if err != nil {
		panic(err)
	}
}
