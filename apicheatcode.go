
****************** RELATIONSHIP TABLES ******

type Client struct {
	gorm.Model
Company string
	Trainings []Training   `json:"trainings" gorm:"foreignkey:Refer"`
					 //`gorm:"foreignKey:UserRefer ;references:Company"`
	}

	type Training struct {
	 gorm.Model
 Refer uint
	 Technology     string
	 Participants   string
	}




CREATE TABLE Client(
stu_id     Varchar(8) NOT NULL PRIMARY KEY,
Company   Varchar(20),
Trainings Varchar(20) FOREIGN KEY REFERENCES Training(Technology),
);

CREATE TABLE Training(
Refer    INT,
Technology  VARCHAR(20) PRIMARY KEY,
Participants VARCHAR(20),
);

CREATE TABLE Usaclient(
	id    INT, 
	Company   Varchar(50)NOT NULL,
	Trainings Varchar(50)NOT NULL,
	PRIMARY KEY (id));

	CREATE TABLE Training(
		id    INT,
		Technology  Varchar(50) NOT NULL,
		Participants Varchar(50)NOT NULL,
		PRIMARY KEY (id)
		FOREIGN KEY (Participants)
		REFERENCES Usaclient(Company)
		)ENGINE=INNODB;
		INDEX (Technology,Participants),
  CONSTRAINT `fk_history_member` FOREIGN KEY (id) 
  REFERENCES Usaclient(Trainings)
		)ENGINE=INNODB;



************** REGISTER /Routing ************

	 v1 := r.Group("/ttd")
 {
	v1.GET("/login",obj.API )
 }


************** INPUT ************

   var  inputs LoginInput

   req.ShouldBind(&inputs)


************* VALIDATION *********


if len(inputs.Passsword) < 9 {
req.JSON(201,"invalid password")
	return nil
}


******************OUTPUT ***********

c.AsciiJSON(200 ,Response)


******************QUERY ************

req.db.Create(&Client{Name: "cerner usa", Trainings: []Training{{Technology: "golang" , Participants:"freshers"}, {Technology: "clojure", Participants:"experienced"}}})


******************* secuirty middleware ***********

