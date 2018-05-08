package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
