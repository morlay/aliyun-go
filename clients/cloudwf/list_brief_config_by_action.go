package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListBriefConfigByActionRequest struct {
	AncestorApgroupId int64  `position:"Query" name:"AncestorApgroupId"`
	Limit             int    `position:"Query" name:"Limit"`
	FuzzySearch       string `position:"Query" name:"FuzzySearch"`
}

func (r ListBriefConfigByActionRequest) Invoke(client *sdk.Client) (response *ListBriefConfigByActionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListBriefConfigByActionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListBriefConfigByAction", "", "")

	resp := struct {
		*responses.BaseResponse
		ListBriefConfigByActionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListBriefConfigByActionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListBriefConfigByActionResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
