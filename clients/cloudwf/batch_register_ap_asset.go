package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BatchRegisterApAssetRequest struct {
	requests.RpcRequest
	JsonData string `position:"Query" name:"JsonData"`
}

func (req *BatchRegisterApAssetRequest) Invoke(client *sdk.Client) (resp *BatchRegisterApAssetResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BatchRegisterApAsset", "", "")
	resp = &BatchRegisterApAssetResponse{}
	err = client.DoAction(req, resp)
	return
}

type BatchRegisterApAssetResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
