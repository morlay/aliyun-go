package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetectVulByIdRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	VulId      int    `position:"Query" name:"VulId"`
}

func (req *DetectVulByIdRequest) Invoke(client *sdk.Client) (resp *DetectVulByIdResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "DetectVulById", "", "")
	resp = &DetectVulByIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetectVulByIdResponse struct {
	responses.BaseResponse
	RequestId string
}
