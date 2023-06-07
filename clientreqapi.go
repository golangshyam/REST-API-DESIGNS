package main 

import (
       	"fmt"
       "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

type  LoginInput  struct { 
       Company string  
       Technology     string
       Participants   string
}

      type Client struct {
         gorm.Model

		 	Technology     string
		 	Participants   string
	 		Id int
	 		Company string
         	Trainings []Training `gorm:"foreignKey:Technology;references:Company"`
         }

         type Training struct {
          gorm.Model
          Technology     string
          Participants   string
         }

type  LoginOutput  struct {
      Status   string      `json:"email"`
}

type   Balaram   interface {

            API(c *gin.Context   )

}

type    COMAPI  struct {
         
             db *gorm.DB

}

       // 3. connect to db
	 // 4. create table 

 func Constuctor()  Balaram  {

	db, err := gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/trainingsdb?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")
	}

	db.AutoMigrate(&Client{} )

	return &COMAPI{ db }

}

 func Inputvalidation (  req   *gin.Context   )  *LoginInput{

       var  inputs LoginInput

        req.ShouldBind(&inputs)

	/*

	if len(inputs.Passsword) < 9 {
			    req.JSON(201,"invalid password")
			    return nil
	}

	*/
  
     return  &inputs

 }

func    APIvalidation (inputs * LoginInput , req   *COMAPI   ) bool   {

         var responses Client

	 fmt.Println( inputs)

       req.db.Create(&Client{Company:"cerner usa", Trainings: []Training{{Technology: "golang" , Participants:"freshers"}, {Technology: "clojure", Participants:"experienced"}}})

        //req.db.Where("email = ? AND passsword = ?",inputs.Email ,inputs.Passsword).Find(&responses)

     //  db.Where("email = ?", inputs.Email).First(&responses)

       if   responses.Company == "" {
	 return false
       }

       return true

}

func  (req   *COMAPI  )  API(c *gin.Context   ){

	// 5. Inside api 
       // above 3 steps   
 
        inputs:=  Inputvalidation(c   )

	if inputs == nil {
	     c.JSON(201,"invalid input")
              return 
	}
	
	fmt.Println( inputs)

 
        resp := APIvalidation(  inputs ,req  )

	if resp == false {
	   c.JSON(201,"invalid email or password ")
              return 
	}

  

      var Response  LoginOutput = LoginOutput{ "trining  registred" }

       c.AsciiJSON(200 ,Response)
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