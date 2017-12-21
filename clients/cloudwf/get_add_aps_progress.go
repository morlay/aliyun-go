package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAddApsProgressRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetAddApsProgressRequest) Invoke(client *sdk.Client) (resp *GetAddApsProgressResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetAddApsProgress", "", "")
	resp = &GetAddApsProgressResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAddApsProgressResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
