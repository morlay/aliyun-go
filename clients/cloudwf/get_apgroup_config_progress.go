package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApgroupConfigProgressRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetApgroupConfigProgressRequest) Invoke(client *sdk.Client) (resp *GetApgroupConfigProgressResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApgroupConfigProgress", "", "")
	resp = &GetApgroupConfigProgressResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetApgroupConfigProgressResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
