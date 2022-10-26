package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var water int = 10
var wind int = 5

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type JSONOutput struct {
	Status Status `json:"status"`
}

func DataBase() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100

	water = rand.Intn(max-min+1) + min
	wind = rand.Intn(max-min+1) + min

	fmt.Println("Water : ", water)
	fmt.Println("Wind  : ", wind)

	statusJSON := &Status{Water: water, Wind: wind}
	JSONOutput := &JSONOutput{Status: *statusJSON}
	b, _ := json.MarshalIndent(JSONOutput, "", "   ")
	_ = ioutil.WriteFile("./static/file.json", b, 0644)

}

func loadHTML() {

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/newvalue", func(rw http.ResponseWriter, r *http.Request) {
		DataBase()
	})
	http.ListenAndServe(":8080", nil)
}

func main() {
	DataBase()

	fmt.Println("load html")
	loadHTML()
}
