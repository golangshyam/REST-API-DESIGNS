package main

import (
		"fmt"
		"log"
		"net/http"
					)

func aap(res http.ResponseWriter, req *http.Request){

	log.Fatal("hello kegriwal")

}

func bjp(res http.ResponseWriter, req *http.Request){

	log.Fatal("hello modi")

}

func main(){


	A1:=func(res http.ResponseWriter, req *http.Request){

		log.Fatal("hello congress")
	}


	http.HandleFunc("/delhi",aap)

	http.HandleFunc("/newdelhi",bjp)

	http.HandleFunc("/dilli",A1)

	http.ListenAndServe(":8080",nil)

	fmt.Println("api approch in go")
}

