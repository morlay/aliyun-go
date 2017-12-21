package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
