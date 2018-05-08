package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
