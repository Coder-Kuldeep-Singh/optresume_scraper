package model

import (
	"fmt"
	"log"

	"github.com/optresume/db"
	re "github.com/optresume/results"
)

func Optresume(data []re.Jobs) {
	db, err := db.InitDB()
	if err != nil {
		log.Println("Connection String Failed")
		return
	}
	defer db.Close()
	_, err = db.Exec(`DROP DATABASE IF EXISTS optresume`)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Create database
	_, err = db.Exec(`CREATE DATABASE optresume`)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Database Created Successfully")

	//Use database
	_, err = db.Exec(`USE optresume`)
	if err != nil {
		fmt.Println(err.Error())
	}

	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS jobs ( id int NOT NULL AUTO_INCREMENT,
		url varchar(255) DEFAULT NULL,
		title varchar(255) DEFAULT NULL,
		company varchar(255) DEFAULT NULL,
		location varchar(255) DEFAULT NULL,
		type varchar(255) DEFAULT NULL,
		description  longtext DEFAULT NULL,
		time2posted varchar(255) DEFAULT NULL,
		PRIMARY KEY(id) )  ENGINE = InnoDB DEFAULT CHARSET = latin1;
		`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Error Table not created!! ", err)
		return
	}
	fmt.Println("Table Created Successfully")

	_, err = db.Exec("ALTER TABLE optresume.jobs CONVERT TO CHARACTER SET utf8")
	if err != nil {
		fmt.Println("Error Table not created!! ", err)
		return
	}

	for _, Eject := range data {
		statement := `INSERT INTO jobs(url,
			title,
			company,
			location,
			type,
			description,
			time2posted) VALUES(?,?,?,?,?,?,?)`
		result, err := db.Exec(statement, Eject.View, Eject.Title, Eject.Company, Eject.Location, Eject.Type, Eject.Description, Eject.PostedTime)
		if err != nil {
			fmt.Println("Error to Execute insert statement!!! ", err)
		}
		fmt.Println(result)
		fmt.Println("Statement Executed")
	}

}
