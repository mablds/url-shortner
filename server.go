package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonTeste struct {
	Author  string `json:"author"`
	Project string `json:"project"`
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	teste := JsonTeste{
		Author:  "mablds",
		Project: "github.com/mablds/url-shortner",
	}

	fmt.Println(r)

	jsonBytes, err := json.Marshal(teste)
	if err != nil {

	}

	w.Write(jsonBytes)
}

func main() {
	http.HandleFunc("/", homePageHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
