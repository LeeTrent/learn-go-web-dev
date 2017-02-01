package main

import (
	"log"
	"net/http"
	"text/template"
)

type menuitem struct {
	Name, Description, Price string
}

type menu struct {
	MenuCategory string
	MenuItems    []menuitem
}

type restaurant []menu

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	m := restaurant{
		menu{
			MenuCategory: "Breakfast",
			MenuItems: []menuitem{
				menuitem{
					Name:        "Vegan Breakfast #1",
					Description: "Vegan Breakfast Description #1",
					Price:       "$4.99",
				},
				menuitem{
					Name:        "Vegan Breakfast #2",
					Description: "Vegan Breakfast Description #2",
					Price:       "$6.99",
				},
				menuitem{
					Name:        "Vegan Breakfast #3",
					Description: "Vegan Breakfast Description #3",
					Price:       "$8.99",
				},
			},
		},
		menu{
			MenuCategory: "Lunch",
			MenuItems: []menuitem{
				menuitem{
					Name:        "Vegan Lunch #1",
					Description: "Vegan Lunch Description #1",
					Price:       "$7.99",
				},
				menuitem{
					Name:        "Vegan Lunch #2",
					Description: "Vegan Lunch Description #2",
					Price:       "$9.99",
				},
				menuitem{
					Name:        "Vegan Lunch #3",
					Description: "Vegan Lunch Description #3",
					Price:       "$11.99",
				},
			},
		},
		menu{
			MenuCategory: "Dinner",
			MenuItems: []menuitem{
				menuitem{
					Name:        "Vegan Dinner #1",
					Description: "Vegan Dinner Description #1",
					Price:       "$9.99",
				},
				menuitem{
					Name:        "Vegan Dinner #2",
					Description: "Vegan Dinner Description #2",
					Price:       "$12.99",
				},
				menuitem{
					Name:        "Vegan Dinner #3",
					Description: "Vegan Dinner Description #3",
					Price:       "$14.99",
				},
			},
		},
	}
	err := tpl.ExecuteTemplate(res, "tpl.gohtml", m)
	if err != nil {
		log.Fatalln(err)
	}
}
