
package main 

import (
	"fmt"
       "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
//	_"gin.SetMode(gin.ReleaseMode)"
)

// Input :             :  email,password 
// Validation    : 
// Query              : one table :  dob , email , password ,address ,firstname ,lastname 
// Output            :   redirect to home page 


type  sevalogin  struct {
	SevaName                string            `form:"sevaname"`
	Date                string            `form:"date"`   
	Nooftickets             int                  `form:"nooftickets"` 
}


type sevauserdetails struct{


		SevaName                string           
		Date                    string          
	   Avail                     int            
	   Price                         string  
	  Allowed                      string
}


// use shyamdb;

// INSERT INTO customer_details (name, email, pincode, proof,mobileno,passsword) VALUES ('venkatesh', 'venkatesh.db@gmail.com', '560075', 'ASUPB5901F','9900367097', '1@venkatesh');
//INSERT INTO iustomerdetails (seva_name, date, avail, price,allowed) VALUES ('suprabatam', '12-1-2023', 20, '100 rs','12');

type  sevalogout  struct {

	SevaName       string       `json:"sevaname"`      
   Date           string          `json:"date"`
	Avail     int                  `json:"avail"`
   Price     string            `json:"price"`
  Allowed 	string           `json:"allowed"`
}


type    COMAPI  struct {
         
             db *gorm.DB

}

       // 3. connect to db
	 // 4. create table 

 func Constuctor()   *COMAPI {

	//gin.SetMode(gin.ReleaseMode)
	//router := gin.New()

	//fmt.Println(router)

	//func init() {

		gin.SetMode(gin.ReleaseMode)

	db, err := gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/shyamdb?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")
	}

	db.AutoMigrate(&sevauserdetails{}  )

	return &COMAPI{ db }


}

 func Inputvalidati (  req   *gin.Context   )  *sevalogin{

       var  inputs sevalogin

        req.Bind(&inputs)

	if len(inputs.SevaName) < 1 {
			    req.JSON(201,"invalid password")
			    return nil
	}
  
     return  &inputs

 }

func    APIvalidati (inputs *sevalogin  , req   *COMAPI   ) bool   {

         var responses sevauserdetails

	 fmt.Println( inputs)

        req.db.Where("seva_name = ? AND date = ?",inputs.SevaName ,inputs.Date).Find(&responses)

   //  db.Where("email = ?", inputs.Email).First(&responses)

       if   responses.SevaName == "" {
	 return false
       }

       return true

}

func  (req   *COMAPI  )  API(c *gin.Context   ){

	// 5. Inside api 
       // above 3 steps   
 
        inputs:=  Inputvalidati(c)

	if inputs == nil {
	     c.JSON(201,"invalid input")
              return 
	}
	
	fmt.Println( inputs)

        resp := APIvalidati(inputs,req)

	if resp == false {
	   c.JSON(201,"invalid sevaname or date ")
              return 
	}

      var Response  sevalogout = sevalogout { "suprabatam","12-1-2023",20,"100 rs","12" }

       c.JSON(200 ,Response)
}



func main(){

	//  1. server 
	//  2. register api to server 

           r := gin.Default()

	   obj:= Constuctor()

     v1 := r.Group("/ttd")
	{
	  v1.GET("/login",obj.API )
	  }
	
	 r.Run(":9090")


}