package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OnoffGroupApRadioRequest struct {
	requests.RpcRequest
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
	Disabled  int   `position:"Query" name:"Disabled"`
}

func (req *OnoffGroupApRadioRequest) Invoke(client *sdk.Client) (resp *OnoffGroupApRadioResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "OnoffGroupApRadio", "", "")
	resp = &OnoffGroupApRadioResponse{}
	err = client.DoAction(req, resp)
	return
}

type OnoffGroupApRadioResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
