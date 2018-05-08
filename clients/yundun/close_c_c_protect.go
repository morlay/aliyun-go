package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CloseCCProtectRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (req *CloseCCProtectRequest) Invoke(client *sdk.Client) (resp *CloseCCProtectResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "CloseCCProtect", "", "")
	resp = &CloseCCProtectResponse{}
	err = client.DoAction(req, resp)
	return
}

type CloseCCProtectResponse struct {
	responses.BaseResponse
	RequestId common.String
}
