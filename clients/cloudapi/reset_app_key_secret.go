package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ResetAppKeySecretRequest struct {
	requests.RpcRequest
	AppKey string `position:"Query" name:"AppKey"`
}

func (req *ResetAppKeySecretRequest) Invoke(client *sdk.Client) (resp *ResetAppKeySecretResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ResetAppKeySecret", "apigateway", "")
	resp = &ResetAppKeySecretResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetAppKeySecretResponse struct {
	responses.BaseResponse
	RequestId common.String
}
