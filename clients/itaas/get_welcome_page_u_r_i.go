package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetWelcomePageURIRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
}

func (r GetWelcomePageURIRequest) Invoke(client *sdk.Client) (response *GetWelcomePageURIResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetWelcomePageURIRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("ITaaS", "2017-05-05", "GetWelcomePageURI", "", "")

	resp := struct {
		*responses.BaseResponse
		GetWelcomePageURIResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetWelcomePageURIResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetWelcomePageURIResponse struct {
	RequestId string
	Data      string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList GetWelcomePageURIErrorMessageList
}

type GetWelcomePageURIErrorMessage struct {
	ErrorMessage string
}

type GetWelcomePageURIErrorMessageList []GetWelcomePageURIErrorMessage

func (list *GetWelcomePageURIErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetWelcomePageURIErrorMessage)
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
