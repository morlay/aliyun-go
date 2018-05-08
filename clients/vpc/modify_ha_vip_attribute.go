package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyHaVipAttributeRequest struct {
	requests.RpcRequest
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyHaVipAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyHaVipAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyHaVipAttribute", "vpc", "")
	resp = &ModifyHaVipAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyHaVipAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
