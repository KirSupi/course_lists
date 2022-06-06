package sms_center

import (
	"encoding/json"
	"fmt"
)

type Phone struct {
	Number int64   `json:"number"`
	SMS    SMSList `json:"sms"`
}

type PhoneListItem struct {
	Value Phone
	next  *PhoneListItem
}

type PhonesList struct {
	Len   int
	first *PhoneListItem
}

func (l *PhonesList) IsEmpty() bool {
	return l.first == nil
}

func (l *PhonesList) Append(p Phone) {
	if l.first == nil {
		l.first = &PhoneListItem{Value: p}
		l.Len = 1
	} else {
		lastItem := l.first
		for ; lastItem.next != nil; lastItem = lastItem.next {
		}
		lastItem.next = &PhoneListItem{Value: p}
		l.Len += 1
	}
}

func (l *PhonesList) FindPhone(phoneNumber int64) *Phone {
	if l.IsEmpty() {
		return nil
	}
	for item := l.first; item != nil; item = item.next {
		if item.Value.Number == phoneNumber {
			return &item.Value
		}
	}
	return nil
}

func (l *PhonesList) MarshalJSON() ([]byte, error) {
	if l.Len == 0 {
		return json.Marshal([]interface{}{})
	}
	slice := make([]Phone, 0, l.Len)
	item := l.first
	for ; item.next != nil; item = item.next {
		slice = append(slice, item.Value)
	}
	slice = append(slice, item.Value)
	fmt.Println("phones: ", slice)
	return json.Marshal(slice)
}
