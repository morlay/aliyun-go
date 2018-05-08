package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
