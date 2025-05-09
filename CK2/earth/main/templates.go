package main

import (
	"github.com/thalesfu/golangutils"
	"html/template"
)

var EmpireTemplate *template.Template
var KingdomTemplate *template.Template
var DukeTemplate *template.Template
var CountyTemplate *template.Template
var BaronyTemplate *template.Template
var EarthTemplate *template.Template

func init() {
	var err error
	EarthTemplate, err = template.New("EarthTemplate.txt").Funcs(template.FuncMap{"FirstUpper": golangutils.FirstUpper}).ParseFiles("Ck2/earth/main/EarthTemplate.txt")
	if err != nil {
		panic(err)
	}

	EmpireTemplate, err = template.New("EmpireTemplate.txt").Funcs(template.FuncMap{"FirstUpper": golangutils.FirstUpper}).ParseFiles("Ck2/earth/main/EmpireTemplate.txt")
	if err != nil {
		panic(err)
	}

	KingdomTemplate, err = template.New("KingdomTemplate.txt").Funcs(template.FuncMap{"FirstUpper": golangutils.FirstUpper}).ParseFiles("Ck2/earth/main/KingdomTemplate.txt")
	if err != nil {
		panic(err)
	}

	DukeTemplate, err = template.New("DukeTemplate.txt").Funcs(template.FuncMap{"FirstUpper": golangutils.FirstUpper}).ParseFiles("Ck2/earth/main/DukeTemplate.txt")
	if err != nil {
		panic(err)
	}

	CountyTemplate, err = template.New("CountyTemplate.txt").Funcs(template.FuncMap{"FirstUpper": golangutils.FirstUpper}).ParseFiles("Ck2/earth/main/CountyTemplate.txt")
	if err != nil {
		panic(err)
	}

	BaronyTemplate, err = template.New("BaronyTemplate.txt").Funcs(template.FuncMap{"FirstUpper": golangutils.FirstUpper}).ParseFiles("Ck2/earth/main/BaronyTemplate.txt")
	if err != nil {
		panic(err)
	}
}
