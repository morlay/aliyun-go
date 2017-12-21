package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DetectVulByIdRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	VulId      int    `position:"Query" name:"VulId"`
}

func (r DetectVulByIdRequest) Invoke(client *sdk.Client) (response *DetectVulByIdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DetectVulByIdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "DetectVulById", "", "")

	resp := struct {
		*responses.BaseResponse
		DetectVulByIdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DetectVulByIdResponse

	err = client.DoAction(&req, &resp)
	return
}

type DetectVulByIdResponse struct {
	RequestId string
}
