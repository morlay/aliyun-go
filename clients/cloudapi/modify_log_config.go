package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyLogConfigRequest struct {
	requests.RpcRequest
	SlsProject  string `position:"Query" name:"SlsProject"`
	SlsLogStore string `position:"Query" name:"SlsLogStore"`
	LogType     string `position:"Query" name:"LogType"`
}

func (req *ModifyLogConfigRequest) Invoke(client *sdk.Client) (resp *ModifyLogConfigResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyLogConfig", "apigateway", "")
	resp = &ModifyLogConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyLogConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
