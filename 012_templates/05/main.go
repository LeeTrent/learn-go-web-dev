package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"text/template"
)

type quote struct {
	Date, Open, High, Low, Close, Volume, AdjClose string
}

var quotes []quote
var tpl *template.Template

func init() {
	quotes = parseFile()
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func parseFile() []quote {
	data, err := os.Open("table.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	reader := csv.NewReader(data)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	recs := make([]quote, 0, len(rows))
	for index, row := range rows {
		if index > 0 {
			q := quote{
				Date:     row[0],
				Open:     row[1],
				High:     row[2],
				Low:      row[3],
				Close:    row[4],
				Volume:   row[5],
				AdjClose: row[6],
			}
			recs = append(recs, q)
		}
	}
	return recs

}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "tpl.gohtml", quotes)
	if err != nil {
		log.Fatalln(err)
	}
}
