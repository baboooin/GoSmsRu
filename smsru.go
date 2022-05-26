package go_sms_ru

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Smsru struct {
	Token string
	Http  *http.Client
}

func SmsRU(token string) (sms *Smsru) {
	return &Smsru{
		Token: token,
		Http:  &http.Client{},
	}
}

func (sms *Smsru) Send(mobile string, message string) (map[string]interface{}, error) {
	URI := fmt.Sprintf("https://sms.ru/sms/send?api_id=%s&to=%s&msg=%s,&json=1", sms.Token, mobile, message)
	res, err := sms.Http.Get(URI)
	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resBytes := []byte(string(resBody))
	var jsonRes map[string]interface{}
	err = json.Unmarshal(resBytes, &jsonRes)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}

func (sms *Smsru) Call(mobile string) (map[string]interface{}, error) {
	URI := fmt.Sprintf("https://sms.ru/code/call?api_id=%s&phone=%s", sms.Token, mobile)
	res, err := sms.Http.Get(URI)
	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resBytes := []byte(string(resBody))
	var jsonRes map[string]interface{}
	err = json.Unmarshal(resBytes, &jsonRes)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}

func (sms *Smsru) Balance() (map[string]interface{}, error) {
	URI := fmt.Sprintf("https://sms.ru/my/balance?api_id=%s&json=1", sms.Token)
	res, err := sms.Http.Get(URI)
	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resBytes := []byte(string(resBody))
	var jsonRes map[string]interface{}
	err = json.Unmarshal(resBytes, &jsonRes)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}

func (sms *Smsru) Limit() (map[string]interface{}, error) {
	URI := fmt.Sprintf("https://sms.ru/my/limit?api_id=%s&json=1", sms.Token)
	res, err := sms.Http.Get(URI)
	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resBytes := []byte(string(resBody))
	var jsonRes map[string]interface{}
	err = json.Unmarshal(resBytes, &jsonRes)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}
