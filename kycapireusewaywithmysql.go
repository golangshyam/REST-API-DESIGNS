package main 

import (
	"fmt"
       "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

type  Product  struct {
	Name     string            
      Price      int               
  	Pincode   string        
	Availableoffer string
}


/*

 INSERT INTO products (name, price, pincode, availableoffer) VALUES ('Iphone13', '50000', '560075', 'cashback 90 bottle 1+2');

INSERT INTO products (name, price, pincode, availableoffer) VALUES ('Iphone14', '50000', '560075', 'cashback 90 bottle 1+2');

INSERT INTO products (name, price, pincode, availableoffer) VALUES ('Iphone15', '99000', '560075', 'cashback 90 bottle 1+2');


*/



type  Review struct {
	Title     string      
       Descripttion     string            
 }

/*

     INSERT INTO reviews (title,descripttion) VALUES ('worthy buy', 'in photo not my face mahesh babu is coming');

    INSERT INTO reviews (title,descripttion) VALUES ('phoen hanging', 'when my wife calls phone is hanging');


*/

type Specification struct{
	  
    Spec   string
}

/*

     INSERT INTO specifications (spec) VALUES ('64mp 7g  128 gb  qualcom processor');

    INSERT INTO specifications (spec) VALUES ('128 mp 10g  1256 gb  qualcom +++++ processor');


*/

var db *gorm.DB


type  Input struct {
	Title     string      `form:"title"`
 }


 func Inputvalidation (  req  *gin.Context  )  *Input{

       var  inputs Input

        req.Bind(&inputs)

	if len(inputs.Title) <0 {
			    req.JSON(201,"invalid name")
			    return nil
	}
  
     return  &inputs

 }

func    APIvalidation (inputs * Input  , req   *gin.Context  ) bool   {

      var responses Product

      db.Where("name = ?", inputs.Title).First(&responses)

       if   responses.Name == "" {
	 return false

       }

       return true

}


func  Productdbquery( inputs * Input  )   *[]Product{

      var response []Product

      db.Where("name = ?", inputs.Title).Find(&response)

      return &response

}

func  Reviewsdbquery( inputs * Input  )   *[]Review{

      var response []Review

      db.Find(&response)

      return &response

}

func  Sepcdbquery( inputs * Input  )   *[]Specification{

      var response []Specification

      db.Find(&response)

      return &response

}


func API( req   *gin.Context   ){

	// 5. Inside api 
       // above 3 steps   
 
        inputs:=  Inputvalidation(req   )

	if inputs == nil {
	     req.JSON(201,"invalid input")
              return 
	}
	
	fmt.Println( inputs)

        resp := APIvalidation(  inputs ,req  )

	if resp == false {
	   req.JSON(201,"product doesnt exist have pray god ")
              return 
	}

	 var finalproduct map[string]interface{} = make (  map[string]interface{} )

          response:=          Productdbquery(inputs   )

        if response ==nil{
          req.JSON(201,"db operation failed ")
              return 
	}

       finalproduct["product"] = response

         response1:=          Reviewsdbquery( inputs  ) 

	 if response1 ==nil{
             req.JSON(201,"db operation failed ")
              return 
	}

	  finalproduct["reviews"] = response1


         response2:=         Sepcdbquery( inputs  )
	 
	 if response2 ==nil{
             req.JSON(201,"db operation failed ")
              return 
	}
    
      finalproduct["spec"] = response2
  
       req.JSON(200,finalproduct)
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
	db.AutoMigrate(&Product{} , &Review{}  , &Specification{} )

}

func main(){

	//  1. server 
	//  2. register api to server 
	 
           r := gin.Default()

     v1 := r.Group("/mobiles")
	{

	  v1.GET("/kart",API )

	  }
	
	 r.Run(":9090")

	}
