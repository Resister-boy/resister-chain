package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./blockchain"
)

const PORT string = ":4000"

type HomeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("template/index.html"))
	data := HomeData("Home", blockchain.getBlockchain().AllBlocks())
	templ.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}