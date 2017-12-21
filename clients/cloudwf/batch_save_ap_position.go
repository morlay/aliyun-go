package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
