package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListConfigByActionRequest struct {
	requests.RpcRequest
	SearchName string `position:"Query" name:"SearchName"`
	Limit      int    `position:"Query" name:"Limit"`
	ActionName string `position:"Query" name:"ActionName"`
}

func (req *ListConfigByActionRequest) Invoke(client *sdk.Client) (resp *ListConfigByActionResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListConfigByAction", "", "")
	resp = &ListConfigByActionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListConfigByActionResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
