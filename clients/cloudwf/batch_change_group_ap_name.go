package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BatchChangeGroupApNameRequest struct {
	requests.RpcRequest
	JsonData string `position:"Query" name:"JsonData"`
}

func (req *BatchChangeGroupApNameRequest) Invoke(client *sdk.Client) (resp *BatchChangeGroupApNameResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BatchChangeGroupApName", "", "")
	resp = &BatchChangeGroupApNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type BatchChangeGroupApNameResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
