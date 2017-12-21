package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApPositionStatusRequest struct {
	requests.RpcRequest
	JsonData string `position:"Query" name:"JsonData"`
}

func (req *ListApPositionStatusRequest) Invoke(client *sdk.Client) (resp *ListApPositionStatusResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApPositionStatus", "", "")
	resp = &ListApPositionStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListApPositionStatusResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
