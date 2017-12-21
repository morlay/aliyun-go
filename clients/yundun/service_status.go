package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ServiceStatusRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *ServiceStatusRequest) Invoke(client *sdk.Client) (resp *ServiceStatusResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "ServiceStatus", "", "")
	resp = &ServiceStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type ServiceStatusResponse struct {
	responses.BaseResponse
	RequestId string
	PortScan  bool
	VulScan   bool
}
