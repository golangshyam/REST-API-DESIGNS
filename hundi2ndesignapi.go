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

var kg *gorm.DB

type  Hundiinput  struct {

	  Trustname                string               `form:"trustname"`
    Amount                     string                `form:"amount"`          
}

type  TrustDetails  struct {


	 Trustname     		string            
    Amount           	int 
    Date                string
    Refernceno  		string
    Donationstatus		 string
}

type  PersnalDetails  struct {


	 Name      		string            
    Age           	int 
    Gender    		string
    Photoid   		string
    ProofNo  		string
}

// use pgbus;

// INSERT INTO Trust_Details (trustname, amount, date, refernceno ,Donationstatus) VALUES ('tirupati', '25000', '6-1-2023', 'ASUPB5901F','successfully');

type  Outputstatus  struct {

	  Status                string              `json:"status"`
}


func enforceJSONHandler( next http.Handler ) http.Handler {

         log.Print("balaram avatar")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

               log.Print("venkatesh db donation 1 cr ")

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

                                log.Print(" venkatesh db no darshan for less payment")
			       }
 

	        

              log.Print("venkatesh db donation continues more giving ends ")

	})
}

func final(w http.ResponseWriter, r *http.Request) {

       log.Print("god bless u venkatesh & jvt family")

              names  := r.URL.Query().Get("name")

             var hundi PeronalDetails 
	     
	     hundi.Name = names

              db.Create(&hundi)



             fmt.Println( "hundi",hundi )

           log.Print("happy sankrathi")

}

	 func init() {

	var err error

	db, err = gorm.Open("mysql", "root:jvt123@tcp(127.0.0.1:3306)/pgbus?charset=utf8&parseTime=True")

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