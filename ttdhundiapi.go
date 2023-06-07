package main 

import (
       	"fmt"
       "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

// Input :             :   Trustname,amount ,  name ,age,gender,photoidproof,no
// Validation    : 
// Query              :  store input to one table : date , donated to , Refernce No, Amount ,donationstatus,
// Output            :  donated sucessfully  


type  Donate  struct {


	  Trustname                string            `form:"trustname"`
          Amount               int                   `form:"amount"`          
}

type  CustomDetails  struct {

	 Trustname     string            
    Amount           int 
    Date                string
    RefernceNo  	string
    Donationstatus 	string
}

type  PersonalDetails  struct {

	 Name      string            
    Age           int 
	Gender    string
       Photoid   string
       ProofNo  string
}



// use pgbus;

// INSERT INTO custom_details (trustname, date, amount, refernceNo,donationstatus) VALUES ('tirumala', '6-1-2023', '10000', 'ASUPB5901F','success');
// INSERT INTO Personal_Details (name, age, gender, photoid,proofNo) VALUES ('shyam', '28', 'male', 'aadhar','943002956401');

type  Donation  struct {

	  Status                string              `json:"status"`
}

type   syam   interface {

            API(c *gin.Context   )

}

type    COMAPI  struct {
         
             db *gorm.DB

}

       // 3. connect to db
	 // 4. create table 

 func Constuctor()  syam  {

	db, err := gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/shyamdb?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")
	}

	db.AutoMigrate(&CustomDetails{},PersonalDetails{})

	return &COMAPI{ db }

}

 func Donatevalidation (  req   *gin.Context   )  *Donate{

       var  inputs Donate 

        req.Bind(&inputs)

	if len(inputs.Trustname) < 1 {

			    req.JSON(201,"invalid trustname")
			    return nil
	}
  
     return  &inputs

 }

func    queryvalid (inputs * Donate , req   *COMAPI   ) bool   {

         var responses CustomDetails

	 fmt.Println( inputs)

        req.db.Where("trustname = ? AND amount = ?",inputs.Trustname ,inputs.Amount).Find(&responses)

   //  db.Where("email = ?", inputs.Email).First(&responses)

       if   responses.Trustname == "" {
	 return false
       }

       return true

}

func  (req   *COMAPI  )  API(c *gin.Context   ){

	// 5. Inside api 
       // above 3 steps   
 
        inputs:= Donatevalidation(c)

	if inputs == nil {
	     c.JSON(201,"invalid input")
              return 
	}
	
	fmt.Println( inputs)

        resp := queryvalid (  inputs ,req  )

	if resp == false {
	   c.JSON(201,"invalid trust or amount ")
              return 
	}

      var Response Donation  = Donation{ Status: "amount transaction sucessfully" }

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