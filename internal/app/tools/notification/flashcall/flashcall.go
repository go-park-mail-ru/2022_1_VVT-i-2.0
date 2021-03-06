package flashcall

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
)

type Flashcaller struct {
	apiKey string
	email  string
}

// TODO: отправлять по хттпС
const (
	flasgcallUrlFmt        = "http://%s:%s@gate.smsaero.ru/v2/flashcall/send?phone=%s&code=%s"
	phoneAlreadyInQueueMsg = "Validation error."
)

// {
// 	"success": true,
// 	"data": {
// 		"id": 1,
// 		"status": 0,
// 		"code": "1234",
// 		"phone": "79990000000",
// 		"cost": "0.59",
// 		"timeCreate": 1646926190,
// 		"timeUpdate": 1646926190
// 	},
// 	"message": null
// }

type respStatus struct {
	Status int `json:"status"`
}

type flashcallResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *respStatus `json:"data"`
}

func NewFlashcaller(cfg *conf.NotificatorConfig) *Flashcaller {
	return &Flashcaller{
		apiKey: cfg.ApiKey,
		email:  cfg.Email,
	}
}

func (f Flashcaller) SendCode(phone string, code string) error {

	client := http.Client{Timeout: 7 * time.Second}

	response, err := client.Get(fmt.Sprintf(flasgcallUrlFmt, f.email, f.apiKey, phone, code))
	if err != nil {
		return servErrors.NewError(servErrors.FLASHCALL_RESPONSE_ERR, "error getting response from flashcall server: "+err.Error())
	}

	bodyBuf, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		return servErrors.NewError(servErrors.FLASHCALL_RESPONSE_ERR, "error reading response body from flashcall server: "+err.Error())
	}
	var respBody flashcallResponse
	if err := json.Unmarshal(bodyBuf, &respBody); err != nil {
		return servErrors.NewError(servErrors.FLASHCALL_RESPONSE_ERR, "error unmarshaling response from flashcall server: "+err.Error())
	}

	if !respBody.Success && respBody.Message == phoneAlreadyInQueueMsg {
		return servErrors.NewError(servErrors.FLASHCALL_PHONE_ALREADY_IN_QUEUE, "")
	}

	if respBody.Data == nil {
		return servErrors.NewError(servErrors.FLASHCALL_RESPONSE_ERR, "flashcall server failed to sent request to queue: "+respBody.Message)
	}

	return nil
}
