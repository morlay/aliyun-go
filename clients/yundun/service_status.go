package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ServiceStatusRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (r ServiceStatusRequest) Invoke(client *sdk.Client) (response *ServiceStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ServiceStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "ServiceStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		ServiceStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ServiceStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type ServiceStatusResponse struct {
	RequestId string
	PortScan  bool
	VulScan   bool
}
