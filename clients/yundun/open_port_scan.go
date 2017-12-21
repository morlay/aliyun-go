package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type OpenPortScanRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r OpenPortScanRequest) Invoke(client *sdk.Client) (response *OpenPortScanResponse, err error) {
	req := struct {
		*requests.RpcRequest
		OpenPortScanRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "OpenPortScan", "", "")

	resp := struct {
		*responses.BaseResponse
		OpenPortScanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.OpenPortScanResponse

	err = client.DoAction(&req, &resp)
	return
}

type OpenPortScanResponse struct {
	RequestId string
}
