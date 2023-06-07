
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


type  LoginInput  struct {
	  Email                string            `form:"email"`
    Passsword      string            `form:"passsword"`          
}

type  CustomerDetails  struct {
	 Name     string            
         Email     string  
  	 Pincode   string  
	Proof  string
	Mobileno string
	Passsword      string      
}


// use pgbus;

// INSERT INTO customer_details (name, email, pincode, proof,mobileno,passsword) VALUES ('venkatesh', 'venkatesh.db@gmail.com', '560075', 'ASUPB5901F','9900367097', '1@venkatesh');


type  LoginOutput  struct {
	  Status                string              `json:"email"`
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

	db.AutoMigrate(&CustomerDetails{}  )

	return &COMAPI{ db }


}

 func Inputvalidation (  req   *gin.Context   )  *LoginInput{

       var  inputs LoginInput

        req.Bind(&inputs)

	if len(inputs.Passsword) < 9 {
			    req.JSON(201,"invalid password")
			    return nil
	}
  
     return  &inputs

 }

func    APIvalidation (inputs * LoginInput , req   *COMAPI   ) bool   {

         var responses CustomerDetails

	 fmt.Println( inputs)

        req.db.Where("email = ? AND passsword = ?",inputs.Email ,inputs.Passsword).Find(&responses)

   //  db.Where("email = ?", inputs.Email).First(&responses)

       if   responses.Name == "" {
	 return false
       }

       return true

}

func  (req   *COMAPI  )  API(c *gin.Context   ){

	// 5. Inside api 
       // above 3 steps   
 
        inputs:=  Inputvalidation(c)

	if inputs == nil {
	     c.JSON(201,"invalid input")
              return 
	}
	
	fmt.Println( inputs)

        resp := APIvalidation(inputs,c)

	if resp == false {
	   c.JSON(201,"invalid email or password ")
              return 
	}

      var Response  LoginOutput = LoginOutput{ "login sucessfully" }

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