package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DetectVulByIpRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	VulIp      string `position:"Query" name:"VulIp"`
}

func (req *DetectVulByIpRequest) Invoke(client *sdk.Client) (resp *DetectVulByIpResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "DetectVulByIp", "", "")
	resp = &DetectVulByIpResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetectVulByIpResponse struct {
	responses.BaseResponse
	RequestId common.String
}
