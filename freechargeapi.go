package main 

import (
	"fmt"
       "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	 _"gorm.io/gorm"
	_ "net/http"
	_ "github.com/gorilla/mux"
)


/*
// first we have to fallow thse steps

1. write input details (mobile number,otp)

2.do the validation part=1) enter the valid 10 digit mobile number
						2) enter the valid 4 digit otp only
						
						

3. do the query part= check the all the stored details compared with given login details

query should cointains
1. firstname
2.last name
3. mobilenumber
4. pincode

4. write the output= response is login successfully
*/


type  FreechargeLogin  struct {


	Mobile      string       `form:"mobile"`     
    Otp   string			`form:"otp"`
}


type UserDetails struct {


	Mobile 	string

	Firstname  string

	Lastname	string

	Pincode		int

	Otp		int


}

//INSERT INTO userdetails (mobile, firstname, lastname,pincode,otp) VALUES ('9177062074', 'shyam', 'varma','505153','9177');

type Result struct{

	Status	string		`json:"status"`
}

var db *gorm.DB

func Inputvalidity (  req  *gin.Context  )  * FreechargeLogin{

	var  inputs FreechargeLogin

	 req.Bind(&inputs)

 if len(inputs.Mobile)<9 {
			 req.JSON(201,"invalid password")
			 return nil
 }

 if len(inputs.Otp)<3 {
	req.JSON(201,"invalid otp")
	return nil
}

  return  &inputs

}

func    APIvalidity (inputs * FreechargeLogin  , req   *gin.Context  ) bool   {

   var responses UserDetails

   db.Where("mobile = ? AND otp = ?" ,inputs.Mobile,inputs.Otp).Find(&responses)

	if   responses.Mobile == "" {

  			return false

	}

	return true

}

func CJI( req   * gin.Context   ){

	// 5. Inside api 
       // above 3 steps   
 
        inputs:= Inputvalidity(req)

	if inputs == nil {


	     req.JSON(201,"invalid details")

              return 
	}
	
	fmt.Println( inputs)

        resp := APIvalidity(  inputs ,req  )

	if resp == false {

	   req.JSON(201,"invalid mobile number or otp ")

              return 
	}

	var res Result = Result{"you have logged into freecharge"}


	req.JSON(200,res)



}

func init() {

	var err error

	db, err = gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/shyamdb?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")

	}
	db.AutoMigrate(&UserDetails{} )

}

func main(){

	//  1. server 
	//  2. register api to server 
	 
           r := gin.Default()

		   v1 := r.Group("/freecharge")
		   {
	   
			 v1.GET("/login",CJI )
	   
			 }

			r.Run(":9090")
			}

