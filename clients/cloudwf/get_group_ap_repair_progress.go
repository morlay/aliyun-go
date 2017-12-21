package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupApRepairProgressRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetGroupApRepairProgressRequest) Invoke(client *sdk.Client) (resp *GetGroupApRepairProgressResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApRepairProgress", "", "")
	resp = &GetGroupApRepairProgressResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetGroupApRepairProgressResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
