package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DelApPositionRequest struct {
	requests.RpcRequest
	ApAssetId int64 `position:"Query" name:"ApAssetId"`
	MapId     int64 `position:"Query" name:"MapId"`
}

func (req *DelApPositionRequest) Invoke(client *sdk.Client) (resp *DelApPositionResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DelApPosition", "", "")
	resp = &DelApPositionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DelApPositionResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
