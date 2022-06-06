package db

import (
	"course/internal/sms_center"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() (*sql.DB, error) {
	return sql.Open("sqlite3", "./database.db")
}

func GetSMSCenterFromDB(db *sql.DB) (sms_center.SMSCenter, error) {
	var SMSCenter sms_center.SMSCenter
	towersRows, err := db.Query(`SELECT tower_number FROM phones GROUP BY tower_number ORDER BY tower_number`)
	if err != nil {
		return SMSCenter, err
	}
	for towersRows.Next() {
		var tower sms_center.Tower
		err = towersRows.Scan(&tower.Number)
		fmt.Println("tower.Number", tower.Number)
		if err != nil {
			return SMSCenter, err
		}
		phonesRows, err := db.Query(`SELECT number FROM phones WHERE tower_number=?`, tower.Number)
		if err != nil {
			return SMSCenter, err
		}
		for phonesRows.Next() {
			var phone sms_center.Phone
			err = phonesRows.Scan(&phone.Number)
			if err != nil {
				return SMSCenter, err
			}
			smsRows, err := db.Query("SELECT `from`, `to`, `text`, `timestamp`, `read` FROM sms WHERE `to`=?", phone.Number)
			if err != nil {
				return SMSCenter, err
			}
			for smsRows.Next() {
				var sms sms_center.SMS
				err = smsRows.Scan(&sms.From, &sms.To, &sms.Text, &sms.Timestamp, &sms.Read)
				fmt.Println(sms)
				if err != nil {
					return SMSCenter, err
				}
				phone.SMS.Append(sms)
			}
			err = smsRows.Close()
			if err != nil {
				return SMSCenter, err
			}
			tower.Phones.Append(phone)
		}
		err = phonesRows.Close()
		if err != nil {
			return SMSCenter, err
		}
		SMSCenter.Towers.Append(tower)
	}
	err = towersRows.Close()
	fmt.Println(SMSCenter)
	fmt.Println(SMSCenter.Towers.Len)
	return SMSCenter, err
}
