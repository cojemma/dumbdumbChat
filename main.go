package main

import (
	chatcore "dumbdumbChat/chatCore"
	"dumbdumbChat/model"
	"dumbdumbChat/wsMessage"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl, _ := template.ParseFiles("./index.html")
	keyData := chatcore.GetKey()
	tmpl.Execute(w, keyData)
}

func setKey(w http.ResponseWriter, r *http.Request) {
	setKeyReq := model.SetKeyRequest{}
	err := json.NewDecoder(r.Body).Decode(&setKeyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chatcore.SetKey(setKeyReq)
}

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", serveHome)

	mux.HandleFunc("/send_value", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("start sending")
		fmt.Fprint(w, `Hello from Golang!"`)
	})

	mux.HandleFunc("/ws", wsMessage.UpgradeConnection)
	mux.HandleFunc("/setKey", setKey)

	// handler := cors.Default().Handler(mux)

	go chatcore.ChatCore()

	fmt.Println("ready to serve on http://localhost:8001/")
	http.ListenAndServe(":8001", mux)
}
