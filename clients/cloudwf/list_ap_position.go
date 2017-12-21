package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApPositionRequest struct {
	requests.RpcRequest
	MapId int64 `position:"Query" name:"MapId"`
}

func (req *ListApPositionRequest) Invoke(client *sdk.Client) (resp *ListApPositionResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApPosition", "", "")
	resp = &ListApPositionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListApPositionResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
