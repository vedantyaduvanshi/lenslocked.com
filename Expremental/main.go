package main

import (
	"html/template"
	"os"
)

type GivingNames struct {
	Name    string
	DogName string
	Nomber  int
	Float   float64
	Slice   []string
	Map     map[string]string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	cons := GivingNames{
		Name:    "Bob barker",
		DogName: "LOduKutta",
		Nomber:  4,
		Float:   10.164,
		Slice:   []string{"a", "b", "c"},
		Map: map[string]string{
			"key1": "value1",
			"key2": "value",
		},
	}

	err = t.Execute(os.Stdout, cons)
	if err != nil {
		panic(err)
	}
}
