package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupApRadioConfigProgressRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetGroupApRadioConfigProgressRequest) Invoke(client *sdk.Client) (resp *GetGroupApRadioConfigProgressResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApRadioConfigProgress", "", "")
	resp = &GetGroupApRadioConfigProgressResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetGroupApRadioConfigProgressResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
