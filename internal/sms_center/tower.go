package sms_center

import (
	"encoding/json"
	"errors"
)

type Tower struct {
	Number int        `json:"number"`
	Phones PhonesList `json:"phones"`
}

type TowersListItem struct {
	Value Tower
	next  *TowersListItem
}

type TowersList struct {
	Len   int
	first *TowersListItem
}

func (l *TowersList) MarshalJSON() ([]byte, error) {
	if l.Len == 0 {
		return json.Marshal([]interface{}{})
	}
	slice := make([]Tower, 0, l.Len)
	item := l.first
	for ; item.next != nil; item = item.next {
		slice = append(slice, item.Value)
	}
	slice = append(slice, item.Value)
	return json.Marshal(slice)
}

func (l *TowersList) Append(t Tower) {
	if l.first == nil {
		l.first = &TowersListItem{Value: t}
		l.Len = 1
	} else {
		lastItem := l.first
		for ; lastItem.next != nil; lastItem = lastItem.next {
		}
		lastItem.next = &TowersListItem{Value: t}
		l.Len += 1
	}
}
func (l *TowersList) Create() {
	newTower := Tower{Number: l.Len + 1}
	l.Append(newTower)
}
func (l *TowersList) FindByNumber(towerNumber int) *Tower {
	if l.first == nil {
		return nil
	}
	for item := l.first; item != nil; item = item.next {
		if item.Value.Number == towerNumber {
			return &item.Value
		}
	}
	return nil
}
func (l *TowersList) FindPhone(phoneNumber int64) (phone *Phone, err error) {
	if l.first == nil {
		return nil, errors.New("list of towers is empty")
	} else {
		for item := l.first; item != nil; item = item.next {
			if phone = item.Value.Phones.FindPhone(phoneNumber); phone != nil {
				return phone, nil
			}
		}
		return nil, errors.New("can't search phone")
	}
}
