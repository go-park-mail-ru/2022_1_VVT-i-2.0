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
	flasgcallUrlFmt = "http://%s:%s@gate.smsaero.ru/v2/flashcall/send?phone=%s&code=%s"
	contentTypeJson = "application/json"
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

type flashcallResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewFlashcaller(cfg *conf.NotificatorConfig) *Flashcaller {
	return &Flashcaller{
		apiKey: cfg.ApiKey,
		email:  cfg.Email,
	}
}

func (f Flashcaller) SendCode(phone string, code string) error {

	client := http.Client{Timeout: 3 * time.Second}

	fmt.Println(fmt.Sprintf(flasgcallUrlFmt, f.email, f.apiKey, phone, code))
	response, err := client.Get(fmt.Sprintf(flasgcallUrlFmt, f.email, f.apiKey, phone, code))
	fmt.Println("-----response----")
	fmt.Print(response)
	fmt.Println("-----end-response----")
	if err != nil {
		return servErrors.NewError(servErrors.FLASHCALL_RESPONSE_ERR, "error getting response from flashcall server: "+err.Error())
	}

	bodyBuf, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	fmt.Println("-----body response----")
	fmt.Print(bodyBuf)
	fmt.Println("-----body end-response----")

	if err != nil {
		return servErrors.NewError(servErrors.FLASHCALL_RESPONSE_ERR, "error reading response body from flashcall server: "+err.Error())
	}
	var respBody flashcallResponse
	if err := json.Unmarshal(bodyBuf, &respBody); err != nil {
		return servErrors.NewError(servErrors.FLASHCALL_RESPONSE_ERR, "error unmarshaling response from flashcall server: "+err.Error())
	}

	if !respBody.Success {
		return servErrors.NewError(servErrors.FLASHCALL_RESPONSE_ERR, "flashcall server failed to sent request to queue: "+respBody.Message)
	}
	return nil
}
