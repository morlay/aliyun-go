package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CloseVulScanRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r CloseVulScanRequest) Invoke(client *sdk.Client) (response *CloseVulScanResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CloseVulScanRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "CloseVulScan", "", "")

	resp := struct {
		*responses.BaseResponse
		CloseVulScanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CloseVulScanResponse

	err = client.DoAction(&req, &resp)
	return
}

type CloseVulScanResponse struct {
	RequestId string
}
