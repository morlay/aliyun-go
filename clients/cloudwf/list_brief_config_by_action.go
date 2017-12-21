package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListBriefConfigByActionRequest struct {
	requests.RpcRequest
	AncestorApgroupId int64  `position:"Query" name:"AncestorApgroupId"`
	Limit             int    `position:"Query" name:"Limit"`
	FuzzySearch       string `position:"Query" name:"FuzzySearch"`
}

func (req *ListBriefConfigByActionRequest) Invoke(client *sdk.Client) (resp *ListBriefConfigByActionResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListBriefConfigByAction", "", "")
	resp = &ListBriefConfigByActionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListBriefConfigByActionResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
