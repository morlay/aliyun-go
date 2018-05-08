package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UnassociateGlobalAccelerationInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	InstanceType                 string `position:"Query" name:"InstanceType"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
}

func (req *UnassociateGlobalAccelerationInstanceRequest) Invoke(client *sdk.Client) (resp *UnassociateGlobalAccelerationInstanceResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "UnassociateGlobalAccelerationInstance", "vpc", "")
	resp = &UnassociateGlobalAccelerationInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnassociateGlobalAccelerationInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
