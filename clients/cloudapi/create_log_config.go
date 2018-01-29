package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateLogConfigRequest struct {
	requests.RpcRequest
	SlsProject  string `position:"Query" name:"SlsProject"`
	SlsLogStore string `position:"Query" name:"SlsLogStore"`
	LogType     string `position:"Query" name:"LogType"`
}

func (req *CreateLogConfigRequest) Invoke(client *sdk.Client) (resp *CreateLogConfigResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateLogConfig", "apigateway", "")
	resp = &CreateLogConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateLogConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
