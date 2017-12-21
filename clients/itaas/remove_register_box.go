package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveRegisterBoxRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Drsessionid string `position:"Query" name:"Drsessionid"`
}

func (req *RemoveRegisterBoxRequest) Invoke(client *sdk.Client) (resp *RemoveRegisterBoxResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "RemoveRegisterBox", "", "")
	resp = &RemoveRegisterBoxResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveRegisterBoxResponse struct {
	responses.BaseResponse
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList RemoveRegisterBoxErrorMessageList
}

type RemoveRegisterBoxErrorMessage struct {
	ErrorMessage string
}

type RemoveRegisterBoxErrorMessageList []RemoveRegisterBoxErrorMessage

func (list *RemoveRegisterBoxErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveRegisterBoxErrorMessage)
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
