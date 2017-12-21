package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OpenVulScanRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r OpenVulScanRequest) Invoke(client *sdk.Client) (response *OpenVulScanResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OpenVulScanRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "OpenVulScan", "", "")

	resp := struct {
		*responses.BaseResponse
		OpenVulScanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OpenVulScanResponse

	err = client.DoAction(&req, &resp)
	return
}

type OpenVulScanResponse struct {
	RequestId string
}
