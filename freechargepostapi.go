package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

type MobileRecharge struct {
	Mobileno string `form:"mobileno"`
	Operator string `form:"operator"`
	Circle   string `form:"circle"`
}

type RechargeDBS struct {
	Mobileno         string
	Operator         string
	Circle           string
	Home             string
	LoansAndMore     string
	FreechargeWallet string
	Offers           string
	Blog             string
	HelpAndSupport   string
	MyTransactions   string
}

//INSERT INTO userdetails (mobileno, operator, circle, home, loansandmore, freechargewallet, offers, blog, helpandsupport, mytransactions) VALUES ('9177062074', 'jio', 'telangana','hitshomepage','eligibleloans', 'walletpage','suitableoffers','articles','customersupport','lasttransactions');

type RechargeResponse struct {
	Message string
}

var cg *gorm.DB

func CON(res http.ResponseWriter, req *http.Request) {

	// 5. Inside api
	// above 3 steps

	rets := mux.Vars(req)

	var inputs MobileRecharge

	id := req.URL.Query().Get("mobileno")

	inputs.Mobileno = rets["mobileno"]
	inputs.Operator = rets["operator"]
	inputs.Circle = rets["circle"]
	//inputs.Address = rets["address"]
	//inputs.Pincode = rets["pincode"]

	if len(inputs.Mobileno) < 7 {

		json.NewEncoder(res).Encode("invalid mobile number")
	}

	if len(inputs.Operator) < 1 {

		json.NewEncoder(res).Encode("invalid oprator selected ")
	}

	var dbinsert RechargeDBS

	dbinsert.Mobileno = id

	fmt.Println(id)

	cg.Create(&dbinsert)

	var resposne RechargeResponse = RechargeResponse{Message: "data saved sucessfully"}

	json.NewEncoder(res).Encode(resposne)

}

// 3. connect to db
// 4. create table

func init() {

	var err error

	cg, err = gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/shyamdb?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")

	}
	cg.AutoMigrate(&RechargeDBS{})

}

func main() {

	//  1. server
	//  2. register api to server

	ret := mux.NewRouter() // Register api to server

	// ret.HandleFunc("/kycs",API).Methods("GET")

	ret.HandleFunc("/rchg", CON).Methods("POST")

	// ret.HandleFunc("/kycs/{name}/{dob}/{pan}/{address}/{pincode}",BJP).Methods("POST")

	http.ListenAndServe(":8080", ret) // server

}
