package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type KycInputc struct {
	Redirecturl string
}

type KycResponsec struct {
	Url string `json:"url"`
}

func MNP(res http.ResponseWriter, req *http.Request) {

	var inputs KycInputc

	id := req.URL.Query().Get("redirecturl")

	fmt.Println(id)

	u, err := url.Parse(id) // http://localhost:8080/session/callback

	if err != nil {
		json.NewEncoder(res).Encode("abhiram jamon")
		return
	}

	if u.Scheme != "http" && u.Host != "localhost" {
		json.NewEncoder(res).Encode("invalid jamon")
		return
	}

	inputs.Redirecturl = id

	var resposne KycResponsec = KycResponsec{Url: "https://www.freecharge"}

	fmt.Println(resposne)

	json.NewEncoder(res).Encode(resposne)

}

func main() {

	//  1. server
	//  2. register api to server

	ret := mux.NewRouter() // Register api to server

	//ret.HandleFunc("/kycs",API).Methods("POST")

	ret.HandleFunc("/kycs", MNP).Methods("GET")

	http.ListenAndServe(":8080", ret) // server

}
