package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetWelcomePageURIRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
}

func (req *GetWelcomePageURIRequest) Invoke(client *sdk.Client) (resp *GetWelcomePageURIResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "GetWelcomePageURI", "", "")
	resp = &GetWelcomePageURIResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetWelcomePageURIResponse struct {
	responses.BaseResponse
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
