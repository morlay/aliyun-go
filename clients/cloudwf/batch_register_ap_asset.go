package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BatchRegisterApAssetRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
}

func (r BatchRegisterApAssetRequest) Invoke(client *sdk.Client) (response *BatchRegisterApAssetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BatchRegisterApAssetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BatchRegisterApAsset", "", "")

	resp := struct {
		*responses.BaseResponse
		BatchRegisterApAssetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BatchRegisterApAssetResponse

	err = client.DoAction(&req, &resp)
	return
}

type BatchRegisterApAssetResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
