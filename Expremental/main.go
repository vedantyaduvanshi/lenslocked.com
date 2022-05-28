package main

import (
	"html/template"
	"os"
)

type GooKhalo struct {
	NaamKyaHai string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := GooKhalo{
		NaamKyaHai: "Bob barker",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
