package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type JsonTeste struct {
	ID        string `json:"id"`
	Link      string `json:"link"`
	RerouteTo string `json:"rerouteLink"`
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	link, ok := r.URL.Query()["link"] // getting query params values, Query() returns a tuple with value and bool

	if !ok || len(link[0]) < 1 {
		id := strings.TrimPrefix(r.URL.Path, "/")
		log.Println(id)
	}

	linkParam := link[0]
	log.Println("Url Param 'link' is: " + string(linkParam))

	var randomID string = randSeq(4)
	linkToRedirect := "localhost:8080/" + randomID

	teste := JsonTeste{
		ID:        string(randomID),
		Link:      string(linkParam),
		RerouteTo: linkToRedirect,
	}

	jsonBytes, err := json.Marshal(teste)
	if err != nil {
		log.Println("Error parsing to JSON using Marshal")
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
