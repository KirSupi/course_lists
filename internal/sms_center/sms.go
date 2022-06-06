package sms_center

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SMS struct {
	From      int64  `json:"from"`
	To        int64  `json:"to"`
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
	Read      bool   `json:"read"`
}

type SMSListItem struct {
	Value SMS
	next  *SMSListItem
}

type SMSList struct {
	Len   int
	first *SMSListItem
}

func (l SMSList) MarshalJSON() ([]byte, error) {
	if l.Len == 0 {
		return json.Marshal([]interface{}{})
	}
	slice := make([]SMS, 0, l.Len)
	item := l.first
	for ; item.next != nil; item = item.next {
		slice = append(slice, item.Value)
	}
	slice = append(slice, item.Value)
	fmt.Println("sms: ", slice)
	return json.Marshal(slice)
}

func (l *SMSList) Get(index int) (sms SMS, err error) {
	if index >= l.Len || index < 0 {
		return sms, errors.New("no index in list")
	}
	item := l.first
	for i := 0; i < index; i++ {
		item = item.next
	}
	return item.Value, nil
}

func (l *SMSList) Append(sms SMS) {
	if l.first == nil {
		l.first = &SMSListItem{Value: sms}
		l.Len = 1
	} else {
		lastItem := l.first
		for ; lastItem.next != nil; lastItem = lastItem.next {
		}
		lastItem.next = &SMSListItem{Value: sms}
		l.Len += 1
	}
}
func (l *SMSList) MarkAllAsRead() {
	for item := l.first; item != nil; item = item.next {
		item.Value.Read = true
	}
}
