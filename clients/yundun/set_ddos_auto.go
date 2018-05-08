package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetDdosAutoRequest struct {
	requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
}

func (req *SetDdosAutoRequest) Invoke(client *sdk.Client) (resp *SetDdosAutoResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "SetDdosAuto", "", "")
	resp = &SetDdosAutoResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDdosAutoResponse struct {
	responses.BaseResponse
	RequestId common.String
}
