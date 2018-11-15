package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var serverId int

func submitHandler(w http.ResponseWriter, r *http.Request) {
	// 解析参数，默认是不会解析的
	r.ParseForm()
	if r.Method == "POST" {
		var user map[string]interface{}
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			serverId += 1
			json.Unmarshal(body, &user)
			fmt.Println("id:", user["id"])
			fmt.Println("email:", user["email"])
			fmt.Println("tel:", user["tel"])
			fmt.Println("serverId:", serverId)

			// response json
			resData := map[string]int{"serverId": serverId}
			resJson, err := json.Marshal(resData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(resJson)
		}
	}
}

func main() {
	serverId = 0
	// handle static files
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	// handle ajax submit requests
	http.HandleFunc("/submit", submitHandler)
	// i dont know if have to handle all invalid routes
 	http.HandleFunc("/unkown", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		//w.Write([]byte("Still developing"))
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
