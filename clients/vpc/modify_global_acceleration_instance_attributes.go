package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyGlobalAccelerationInstanceAttributesRequest struct {
	requests.RpcRequest
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	Name                         string `position:"Query" name:"Name"`
	Description                  string `position:"Query" name:"Description"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
}

func (req *ModifyGlobalAccelerationInstanceAttributesRequest) Invoke(client *sdk.Client) (resp *ModifyGlobalAccelerationInstanceAttributesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyGlobalAccelerationInstanceAttributes", "vpc", "")
	resp = &ModifyGlobalAccelerationInstanceAttributesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyGlobalAccelerationInstanceAttributesResponse struct {
	responses.BaseResponse
	RequestId common.String
}
