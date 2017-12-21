package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GroupDetailsRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *GroupDetailsRequest) Invoke(client *sdk.Client) (resp *GroupDetailsResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GroupDetails", "", "")
	resp = &GroupDetailsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GroupDetailsResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
