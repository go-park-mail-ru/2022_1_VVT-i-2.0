package sms

// import (
// 	"bytes"
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"time"

// 	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
// 	"github.com/pkg/errors"
// )

// type SmsManager struct {
// 	apiKey    string
// 	apiSecret string
// 	brandName string
// }

// type sendSmsRequest struct {
// 	From      string `json:"from"`
// 	Text      string `json:"text"`
// 	Encoding  string `json:"type"`
// 	To        string `json:"to"`
// 	ApiKey    string `json:"api_key"`
// 	ApiSecret string `json:"api_secret"`
// }

// // {"message-count":"1","messages":[{"status":"2","error-text":"Missing api_key"}]}

// type sendSmsResponse struct {
// 	MsgCount int `json:"message-count,string"`
// 	Messages []struct {
// 		Status   int    `json:"status,string"`
// 		ErrorMsg string `json:"error-text"`
// 	}
// }

// func NewSmsManager(cfg *conf.NotificatorConfig) *SmsManager {
// 	return &SmsManager{
// 		apiKey:    cfg.ApiKey,
// 		apiSecret: cfg.ApiSecret,
// 		brandName: cfg.BrandName,
// 	}
// }

// /*
// curl -X "POST" "https://rest.nexmo.com/sms/json"
// 	-d "from=Vonage APIs"
// 	-d "text=A text message sent using the Vonage SMS API"
// 	-d "to=79015020456"
// 	-d "api_key=c4624db8"
// 	-d "api_secret=1SQlhK6gbfJP43Kz"
// */

// const (
// 	sendSmsUrl      = "https://rest.nexmo.com/sms/json"
// 	contentTypeJson = "application/json"
// )

// func (sender *SmsManager) Send(phone string, msg string, encoding string) error {

// 	client := http.Client{Timeout: 3 * time.Second}

// 	reqJsonBody, err := json.Marshal(sendSmsRequest{
// 		From:      sender.brandName,
// 		Text:      msg,
// 		To:        phone,
// 		Encoding:  encoding,
// 		ApiKey:    sender.apiKey,
// 		ApiSecret: sender.apiSecret,
// 	})

// 	if err != nil {
// 		return errors.Wrap(err, "error marshaling to json")
// 	}

// 	requestBody := bytes.NewBuffer(reqJsonBody)

// 	response, err := client.Post(sendSmsUrl, contentTypeJson, requestBody)
// 	// fmt.Println("-----response----")
// 	// fmt.Print(response)
// 	// fmt.Println("-----end-response----")
// 	if err != nil {
// 		return errors.Wrap(err, "error getting response from sms-sender server")
// 	}

// 	bodyBuf, err := ioutil.ReadAll(response.Body)
// 	response.Body.Close()
// 	if err != nil {
// 		return errors.Wrap(err, "error of reading response body from sms-sender service")
// 	}
// 	var respBody sendSmsResponse
// 	if err := json.Unmarshal(bodyBuf, &respBody); err != nil {
// 		return errors.Wrap(err, "error unmarshaling response from sms-sender service")
// 	}

// 	if respBody.MsgCount != 1 {
// 		return servErrors.NewError(servErrors.SENDING_AUTH_CODE, "request to sms-sender server return unexpected body (msg-count!=1)")
// 	}
// 	if respBody.Messages[0].Status != 0 {
// 		return servErrors.NewError(servErrors.SENDING_AUTH_CODE, "request to sms-sender server return error"+respBody.Messages[0].ErrorMsg)
// 	}

// 	return nil
// }

//поднять сервер, который будет ловить колбэки
// ловить отчет об отправке
// если не вышло пробовать отправить еще раз
// логировать
