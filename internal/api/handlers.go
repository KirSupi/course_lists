package api

import (
	"course/internal/sms_center"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func handlerPing(c *gin.Context) {
	c.String(200, "pong")
}
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
func (api *API) handlerGetSMSCenter(c *gin.Context) {
	res, err := PrettyStruct(api.SMSCenter)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	c.String(http.StatusOK, res)
}
func (api *API) handlerCreateTower(c *gin.Context) {
	api.SMSCenter.Towers.Create()
	c.String(http.StatusOK, "ok")
}
func (api *API) handlerCreatePhone(c *gin.Context) {
	var input sms_center.Phone
	err := c.BindJSON(&input)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	towerNumberParam := c.Param("towerNumber")
	towerNumber, err := strconv.Atoi(towerNumberParam)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = api.SMSCenter.CreatePhone(towerNumber, input.Number)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "ok")
}
func (api *API) handlerSendSMS(c *gin.Context) {
	var input sms_center.SMS
	err := c.BindJSON(&input)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = api.SMSCenter.SendSMS(input)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "ok")
}
func (api *API) handlerGetSMSList(c *gin.Context) {
	towerNumberParam := c.Param("towerNumber")
	towerNumber, err := strconv.Atoi(towerNumberParam)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	phoneNumberParam := c.Param("phoneNumber")
	phoneNumber, err := strconv.ParseInt(phoneNumberParam, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	phone, err := api.SMSCenter.GetPhone(towerNumber, phoneNumber)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	res, err := PrettyStruct(phone.SMS)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	c.String(http.StatusOK, res)
}
