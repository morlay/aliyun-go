package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type BatchSaveApPositionRequest struct {
	requests.RpcRequest
	JsonData string `position:"Query" name:"JsonData"`
}

func (req *BatchSaveApPositionRequest) Invoke(client *sdk.Client) (resp *BatchSaveApPositionResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BatchSaveApPosition", "", "")
	resp = &BatchSaveApPositionResponse{}
	err = client.DoAction(req, resp)
	return
}

type BatchSaveApPositionResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
