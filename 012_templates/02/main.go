package main

import (
	"log"
	"net/http"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
}

type regionHotel struct {
	Region string
	Hotels []hotel
}

// RegionHotels blah blah blah
type RegionHotels []regionHotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	rhs := RegionHotels{
		regionHotel{
			Region: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "Northern Hotel 1",
					Address: "44 Monica Blvd",
					City:    "San Francisco",
					Zip:     "94444",
				},
				hotel{
					Name:    "Northern Hotel 2",
					Address: "111 Aariv Blvd",
					City:    "Sacramento",
					Zip:     "94444",
				},
			},
		},
		regionHotel{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Central Hotel 1",
					Address: "1960 Linda Blvd",
					City:    "Fresno",
					Zip:     "95555",
				},
				hotel{
					Name:    "Central Hotel 2",
					Address: "77 Casey Lane",
					City:    "Fresno",
					Zip:     "95555",
				},
			},
		},
		regionHotel{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Southern Hotel 1",
					Address: "1972 Penny Lane",
					City:    "Los Angelos",
					Zip:     "97777",
				},
				hotel{
					Name:    "Southern Hotel 2",
					Address: "1968 Prince Blvd",
					City:    "San Diego",
					Zip:     "97777",
				},
			},
		},
	}
	err := tpl.ExecuteTemplate(res, "tpl.gohtml", rhs)
	if err != nil {
		log.Fatalln(err)
	}
}
