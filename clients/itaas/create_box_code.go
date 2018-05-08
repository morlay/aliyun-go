package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateBoxCodeRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
}

func (req *CreateBoxCodeRequest) Invoke(client *sdk.Client) (resp *CreateBoxCodeResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "CreateBoxCode", "", "")
	resp = &CreateBoxCodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateBoxCodeResponse struct {
	responses.BaseResponse
	RequestId common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
	Success   bool
	ErrorList CreateBoxCodeErrorMessageList
	Data      CreateBoxCodeData
}

type CreateBoxCodeErrorMessage struct {
	ErrorMessage common.String
}

type CreateBoxCodeData struct {
	ClientAppid       common.String
	Code              common.String
	CurrentTimeMillis common.Long
	ExpiresIn         common.Integer
	ExpiresInUnit     common.String
}

type CreateBoxCodeErrorMessageList []CreateBoxCodeErrorMessage

func (list *CreateBoxCodeErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateBoxCodeErrorMessage)
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
