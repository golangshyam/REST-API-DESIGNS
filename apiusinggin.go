package main 

import (
	"fmt"
       "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	 _"gorm.io/gorm"
)


/*
// first we have to fallow thse steps

1. write input details

2.do the validation part=1) registred email id with @
						2) password contains not less than 10 letters
						3)password should combination of uppercase and lowercase and numbers and symbols
						

3. do the query part= check the all the stored details compared with given login details

query should cointains
1. firstname
2.last name
3. mobilenumber
4. dob
5. address
6. city
7. country
8. state
9. pincode

4. write the output= response is login successfully
*/


type  Facebooklogin  struct {


	Email      string       `form:"email"`     
    Password   string			`form:"password"`
	Mobilenumber	string		`form:"mobilenumber"`
}


type Facebooksignup struct{

	Surname 		string
	Email      string            
    Password   string
	Firstname  string
	Lastname	string
	Mobilenumber	string
	Dob			string
	


}

//INSERT INTO query_details (email, password, firstname, lastname,mobilenumber,dob,address,city,country,state,pincode) VALUES ('ssk0041@gmail.com', 'shyamvarma', 'shyam', 'varma','9177062074','04-05-1994','hyderabad','amerpet','india','telangana','505153');

type Facebookhomepage struct{

	Response	string		`json:"response"`
}

var db *gorm.DB

/*
func Inputdetailsvalid (  req  *gin.Context  )  *Facebooklogin{

	var  inputs Facebooklogin

	 req.Bind(&inputs)

 if len(inputs.Password)<1 {

			 req.JSON(201,"invalid password")
			 return nil
 }

  return  &inputs

}
*/
func    Facebooksignupvalid (inputs * Facebooklogin  , req   *gin.Context  ) bool   {

   var responses Facebooksignup

   db.Where("email = ? AND password = ?,mobilenumbe r= ?",inputs.Email,inputs.Password,inputs.Mobilenumber).Find(&responses)

	if   responses.Email == "" {

  			return false

	}

	if   responses.Mobilenumber == "" {

		return false

}
if   responses.Password == "" {

	return false

}

	return true

}

func API( req   *gin.Context   ) *Facebooklogin {

	// 5. Inside api 
       // above 3 steps   
 
       // inputs:= Inputvalid(req   )

		//func Inputdetailsvalid (  req  *gin.Context  )  *Facebooklogin{

			var  inputs Facebooklogin
		
			 req.Bind(&inputs)
		
		 if len(inputs.Password)<1 {
		
					 req.JSON(201,"invalid password")
					 return nil
		 }
		
		 // return  &inputs
		
		

			if len(inputs.Mobilenumber)<1 {

	     req.JSON(201,"invalid mobile number")

              return nil
	}
	
	if len(inputs.Email)<1 {

		req.JSON(201,"invalid email")

			 return nil
	}
	fmt.Println( inputs)

       // resp := Facebooksignupvalid(inputs,req )

	   var resp Facebooksignup

	if resp == false {
	   req.JSON(201,"invalid username or password ")
              return 
	}

	var res Outputres = Outputres{"loginsuccessful"}
	req.JSON(200,res)

	return  &inputs
	
}

func init() {

	var err error

	db, err = gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/shyamdb?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")

	}
	db.AutoMigrate(&Querydetails{} )

}

func main(){

	//  1. server 
	//  2. register api to server 
	 
           r := gin.Default()

     v1 := r.Group("/ttd")
	{

	  v1.GET("/login",API )

	  }
	
	 r.Run(":9090")


}