package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetectVulByIpRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	VulIp      string `position:"Query" name:"VulIp"`
}

func (r DetectVulByIpRequest) Invoke(client *sdk.Client) (response *DetectVulByIpResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DetectVulByIpRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "DetectVulByIp", "", "")

	resp := struct {
		*responses.BaseResponse
		DetectVulByIpResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DetectVulByIpResponse

	err = client.DoAction(&req, &resp)
	return
}

type DetectVulByIpResponse struct {
	RequestId string
}
