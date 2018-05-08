package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AssociateHaVipRequest struct {
	requests.RpcRequest
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AssociateHaVipRequest) Invoke(client *sdk.Client) (resp *AssociateHaVipResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "AssociateHaVip", "vpc", "")
	resp = &AssociateHaVipResponse{}
	err = client.DoAction(req, resp)
	return
}

type AssociateHaVipResponse struct {
	responses.BaseResponse
	RequestId common.String
}
