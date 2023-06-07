package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

// Input :                      :     seva name  : suprbhtam ,  date
// Validation              :
// Query                     :     one table :   seva name  : suprbhtam ,  date
// Query Output     :       available : 9 ,  price ,:50 ,  person allowed 1
// API output           :        above same

type LoginInputs struct {
	SevaName    string `form:"sevaname"`
	Date        string `form:"date"`
	Nooftickets int    `form:"nooftickets"`
}

type Iustomerdetails struct {
	SevaName string
	Date     string
	Avail    int
	Price    string
	Allowed  string
}

//use pgbus;

//INSERT INTO iustomerdetails (seva_name, date, avail, price,allowed) VALUES ('archana', '2023-1-2', 20, '50','1 person');

// INSERT INTO iustomerdetails (seva_name, date, avail, price,allowed) VALUES ('archana', '2023-1-2', 20, '50','1 person');

type LoginOutputs struct {
	SevaName string `json:"sevaname"`
	Date     string `json:"date"`
	Avail    int    `json:"avail"`
	Price    string `json:"price"`
	Allowed  string `json:"allowed"`
}

var db *gorm.DB

func Inputvalidations(req *gin.Context) *LoginInputs {

	var inputs LoginInputs

	req.Bind(&inputs)

	if len(inputs.SevaName) < 9 {
		req.JSON(201, "invalid password")
		return nil
	}

	return &inputs

}

func APIvalidations(inputs *LoginInputs, req *gin.Context) bool {

	var responses Iustomerdetails

	fmt.Println(inputs)

	db.Where("seva_name = ? AND date = ?  AND  avail =? ", inputs.SevaName, inputs.Date, inputs.Nooftickets).Find(&responses)

	//  db.Where("email = ?", inputs.Email).First(&responses)

	if responses.SevaName == "" {
		return false
	}

	return true

}

func APIP(req *gin.Context) {

	// 5. Inside api
	// above 3 steps

	inputs := Inputvalidations(req)

	if inputs == nil {
		req.JSON(201, "invalid input")
		return
	}

	fmt.Println(inputs)

	resp := APIvalidations(inputs, req)

	if resp == false {
		req.JSON(201, "invalid email or password ")
		return
	}

	var Response LoginOutputs = LoginOutputs{"archana", "may 2023", 12, "50 rs", "1 person allowed"}

	req.JSON(200, Response)
}

// 3. connect to db
// 4. create table

func init() {

	var err error

	db, err = gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/shyamdb?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")

	}
	db.AutoMigrate(&Iustomerdetail{})

}

func main() {

	//  1. server
	//  2. register api to server

	r := gin.Default()

	v1 := r.Group("/ttd")
	{

		v1.GET("/login", APIP)

	}

	r.Run(":9090")

}
