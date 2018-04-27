package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddGlobalAccelerationInstanceIpRequest struct {
	requests.RpcRequest
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	IpInstanceId                 string `position:"Query" name:"IpInstanceId"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
}

func (req *AddGlobalAccelerationInstanceIpRequest) Invoke(client *sdk.Client) (resp *AddGlobalAccelerationInstanceIpResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "AddGlobalAccelerationInstanceIp", "vpc", "")
	resp = &AddGlobalAccelerationInstanceIpResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddGlobalAccelerationInstanceIpResponse struct {
	responses.BaseResponse
	RequestId string
}
