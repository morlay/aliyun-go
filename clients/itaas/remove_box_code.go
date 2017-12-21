package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveBoxCodeRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Code        string `position:"Query" name:"Code"`
}

func (req *RemoveBoxCodeRequest) Invoke(client *sdk.Client) (resp *RemoveBoxCodeResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "RemoveBoxCode", "", "")
	resp = &RemoveBoxCodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveBoxCodeResponse struct {
	responses.BaseResponse
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList RemoveBoxCodeErrorMessageList
}

type RemoveBoxCodeErrorMessage struct {
	ErrorMessage string
}

type RemoveBoxCodeErrorMessageList []RemoveBoxCodeErrorMessage

func (list *RemoveBoxCodeErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveBoxCodeErrorMessage)
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
