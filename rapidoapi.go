// captain sign up details

//Input

// fullname
// city
// findus
// validdrivingliecence
// mobileno
// otp

//captainsignupdeatils

// fullname
// city
// findus
// validdrivingliecence
// mobileno
// otp

// output 

// Congrats!
// Thank you for showing interest.
// We will contact you soon.



package main

import (
	 "log"
	"net/http"
	// "strconv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

var db *gorm.DB

type  CaptainLogin  struct {
	Mobileno	string            `form:"mobileno"`
        Otp		string             `form:"otp"`          
}

type  CaptainDetails  struct {
	Fullname string
	City	string
	Findus	string
	Validdrivingliecence string
	Mobileno string
	Otp	string
}

type  LoginOutput  struct {

	  Status                string              `json:"status"`
}


func enforceJSONHandler( next http.Handler ) http.Handler {

         log.Print("Running Signup API")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

         log.Print(" Captain Entering Details ")

          var inputs CaptainLogin

          inputs.Mobileno    = r.URL.Query().Get("mobileno")

          inputs.Otp   = r.URL.Query().Get("otp")
 	 
		
	  if len(inputs.Mobileno) < 10 {

			log.Print("Invalid Mobile Number")

			return

	  }

          if inputs.Otp == "1234" {

          next.ServeHTTP(w, r)

          var Rider CaptainDetails = CaptainDetails{Mobileno: inputs.Mobileno, Otp: inputs.Otp}

		  var res LoginOutput=LoginOutput{Status:"rapido login successfully"}

		 // r.JSON(200 ,res)

		 fmt.Println(res)
	  
	  db.Create(&Rider)

	   } else {

          log.Print("invalid OTP")

	  		return
		 		}

              log.Print("Captain Signup Comapleted")

	})
}

func HomepageAPI(w http.ResponseWriter, r *http.Request) {

       log.Print("WELCOME to Rapido")

       log.Print("First User Get 50% off Enjoy..")

}

func init() {

	var err error

	db, err = gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/shyamdb?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")

	}
	db.AutoMigrate(&CaptainDetails{})

}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(HomepageAPI)
	mux.Handle("/signup", enforceJSONHandler(finalHandler))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}	