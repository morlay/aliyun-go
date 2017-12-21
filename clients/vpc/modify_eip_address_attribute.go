package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyEipAddressAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyEipAddressAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyEipAddressAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyEipAddressAttribute", "vpc", "")
	resp = &ModifyEipAddressAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyEipAddressAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
