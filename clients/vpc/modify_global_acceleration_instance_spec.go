package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyGlobalAccelerationInstanceSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                    string `position:"Query" name:"Bandwidth"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
}

func (req *ModifyGlobalAccelerationInstanceSpecRequest) Invoke(client *sdk.Client) (resp *ModifyGlobalAccelerationInstanceSpecResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyGlobalAccelerationInstanceSpec", "vpc", "")
	resp = &ModifyGlobalAccelerationInstanceSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyGlobalAccelerationInstanceSpecResponse struct {
	responses.BaseResponse
	RequestId string
}
