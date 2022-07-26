package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var URL string = "127.0.0.1:8080"

type sampleStruct struct {
	// Name string `json:"abc"`
	ParamA int    `json:"paramA"`
	ParamB string `json:"paramB"`
	ParamC int    `json:"paramC"`
}

func requestHandler(rw http.ResponseWriter, r *http.Request) {
	// Get Request on CURL:  curl "http://127.0.0.1:8080/?paramA=1&paramB=abcd&paramC=-1"
	// Post Request on CURL curl -X POST http://127.0.0.1:8080/ -d '{"paramA": 100, "paramB":"qwert", "paramC":11}'

	if r.Method == "GET" {

		fmt.Println("New Get Request Detected, query data is as follows:")
		fmt.Fprintln(rw, "Get Request Detected\n, Now Echoing:")
		var data sampleStruct
		for k, v := range r.URL.Query() {
			switch k {
			case "paramA":
				data.ParamA, _ = strconv.Atoi(v[0])
			case "paramB":
				data.ParamB = v[0]
			case "paramC":
				data.ParamC, _ = strconv.Atoi(v[0])
			}
		}

		marshall_Struct, err := json.Marshal(data)

		if err != nil {
			log.Fatal(err)
		}
		stringStruct := string(marshall_Struct)
		fmt.Fprintln(rw, stringStruct)
		fmt.Println(stringStruct)

	} else if r.Method == "POST" {
		fmt.Println("New Get Request Detected, query data is as follows:")
		fmt.Fprintln(rw, "Get Request Detected\n, Now Echoing:")

		decoder := json.NewDecoder(r.Body)
		var data sampleStruct
		err := decoder.Decode(&data)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(data)
		fmt.Fprintln(rw, data)
	}

}

func main() {
	fmt.Println("Starting Server at", URL)
	http.HandleFunc("/", requestHandler)
	if err := http.ListenAndServe(URL, nil); err != nil {
		log.Fatal(err)
	}
}
