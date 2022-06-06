package main

import (
	API "course/internal/api"
	DB "course/internal/db"
	"fmt"
)

func main() {
	db, err := DB.GetConnection()
	if err != nil {
		fmt.Println("1: ", err.Error())
	}
	smsCenter, err := DB.GetSMSCenterFromDB(db)
	if err != nil {
		fmt.Println("2: ", err.Error())
	}
	api := API.API{SMSCenter: &smsCenter}
	api.Run()
}
