package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyAppRequest struct {
	requests.RpcRequest
	AppId       int64  `position:"Query" name:"AppId"`
	AppName     string `position:"Query" name:"AppName"`
	Description string `position:"Query" name:"Description"`
}

func (req *ModifyAppRequest) Invoke(client *sdk.Client) (resp *ModifyAppResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyApp", "apigateway", "")
	resp = &ModifyAppResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyAppResponse struct {
	responses.BaseResponse
	RequestId string
}
