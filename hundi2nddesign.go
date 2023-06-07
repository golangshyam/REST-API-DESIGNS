package main

import (
	 "log"
	"net/http"
	"strconv"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

var db *gorm.DB

type  LoginInput  struct {

	  Trustname       string               `form:"trustname"`
    Amount        string                   `form:"amount"`          
}

type  TrustDetails  struct {

	 Trustname     string            
    	Amount           int 
        Date          string
       RefernceNo  		string
       Donationstatus 	string
}

type  PeronalDetails  struct {


	 Name      string            
        Age           int 
        Gender    string
       Photoid   string
       ProofNo  string
}

// use pgbus;

// INSERT INTO customer_details (name, email, pincode, proof,mobileno,passsword) VALUES ('venkatesh', 'venkatesh.db@gmail.com', '560075', 'ASUPB5901F','9900367097', '1@venkatesh');

type  LoginOutput  struct {

	  Status        string              `json:"email"`
}


func enforceJSONHandler( next http.Handler ) http.Handler {

         log.Print("hello shyam govinda govinda...")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

     log.Print("welcome to ttd donate some amount")

         var inputs LoginInput
        inputs.Trustname    = r.URL.Query().Get("trustname")
        inputs.Amount   = r.URL.Query().Get("amount")
 	 
                  
        amounts ,_:=  strconv.Atoi(inputs.Amount)

		fmt.Println( "inputs",inputs )

        if amounts >=10000 {
                next.ServeHTTP(w, r)
                                            
	 var srivanitrust TrustDetails = TrustDetails{Trustname: inputs.Trustname, Amount:amounts  }

		  db.Create(&srivanitrust)

		 } else {

            log.Print(" payment amount is very less")
			       
		}
 
        log.Print("donation process ends ")

	})
}

func final(w http.ResponseWriter, r *http.Request) {

       log.Print("god bless u you and your family")

              names  := r.URL.Query().Get("name")

             var hundi PeronalDetails 
	     
	     hundi.Name = names

              db.Create(&hundi)



             fmt.Println( "hundi",hundi )

           log.Print("thank you visit again ttd")

}

	 func init() {

	var err error

	db, err = gorm.Open("mysql", "root:shyamvarma@tcp(127.0.0.1:3306)/shyamdb?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")

	}
	db.AutoMigrate(&TrustDetails{},&PeronalDetails{} )

}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/hundi", enforceJSONHandler(finalHandler))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}	