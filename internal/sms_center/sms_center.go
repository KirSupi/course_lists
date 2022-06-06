package sms_center

import (
	"errors"
	"time"
)

type SMSCenter struct {
	Towers TowersList `json:"towers"`
}

func (c *SMSCenter) SendSMS(input SMS) error {
	phone, err := c.Towers.FindPhone(input.To)
	if err != nil || phone == nil {
		return errors.New("Телефон не найден")
	}
	input.Timestamp = time.Now().Unix()
	input.Read = false
	phone.SMS.Append(input)
	return nil
}
func (c *SMSCenter) CreatePhone(towerNumber int, phoneNumber int64) error {
	phone, err := c.Towers.FindPhone(phoneNumber)
	if err == nil && phone != nil {
		return errors.New("Этот номер телефона уже существует")
	}
	tower := c.Towers.FindByNumber(towerNumber)
	if tower == nil {
		return errors.New("Вышка не найдена")
	}
	tower.Phones.Append(Phone{Number: phoneNumber})
	return nil
}
func (c *SMSCenter) GetPhone(towerNumber int, phoneNumber int64) (phone Phone, err error) {
	if towerNumber == 0 {
		return phone, errors.New("bad towerNumber")
	}
	if phoneNumber == 0 {
		return phone, errors.New("bad phoneNumber")
	}
	tower := c.Towers.FindByNumber(towerNumber)
	if tower == nil {
		return phone, errors.New("Вышка не найдена")
	}
	phonePtr := tower.Phones.FindPhone(phoneNumber)
	if phonePtr == nil {
		return phone, errors.New("Телефон не найден")
	}
	phone = *phonePtr
	phonePtr.SMS.MarkAllAsRead()
	return
}
