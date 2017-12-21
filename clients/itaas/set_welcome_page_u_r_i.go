package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetWelcomePageURIRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Welcomeuri  string `position:"Query" name:"Welcomeuri"`
}

func (r SetWelcomePageURIRequest) Invoke(client *sdk.Client) (response *SetWelcomePageURIResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetWelcomePageURIRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("ITaaS", "2017-05-05", "SetWelcomePageURI", "", "")

	resp := struct {
		*responses.BaseResponse
		SetWelcomePageURIResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetWelcomePageURIResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetWelcomePageURIResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList SetWelcomePageURIErrorMessageList
}

type SetWelcomePageURIErrorMessage struct {
	ErrorMessage string
}

type SetWelcomePageURIErrorMessageList []SetWelcomePageURIErrorMessage

func (list *SetWelcomePageURIErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SetWelcomePageURIErrorMessage)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
