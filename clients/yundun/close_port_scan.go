package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ClosePortScanRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r ClosePortScanRequest) Invoke(client *sdk.Client) (response *ClosePortScanResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ClosePortScanRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "ClosePortScan", "", "")

	resp := struct {
		*responses.BaseResponse
		ClosePortScanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ClosePortScanResponse

	err = client.DoAction(&req, &resp)
	return
}

type ClosePortScanResponse struct {
	RequestId string
}
