package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UnassociateHaVipRequest struct {
	requests.RpcRequest
	HaVipId              string `position:"Query" name:"HaVipId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UnassociateHaVipRequest) Invoke(client *sdk.Client) (resp *UnassociateHaVipResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "UnassociateHaVip", "vpc", "")
	resp = &UnassociateHaVipResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnassociateHaVipResponse struct {
	responses.BaseResponse
	RequestId string
}
