package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAppRequest struct {
	requests.RpcRequest
	AppName     string `position:"Query" name:"AppName"`
	Description string `position:"Query" name:"Description"`
}

func (req *CreateAppRequest) Invoke(client *sdk.Client) (resp *CreateAppResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateApp", "apigateway", "")
	resp = &CreateAppResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAppResponse struct {
	responses.BaseResponse
	RequestId string
	AppId     int64
}
