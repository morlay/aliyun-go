package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	FromApp              string `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceClass        string `position:"Query" name:"InstanceClass"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ForceUpgrade         string `position:"Query" name:"ForceUpgrade"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
}

func (req *ModifyInstanceSpecRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceSpecResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceSpec", "redisa", "")
	resp = &ModifyInstanceSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceSpecResponse struct {
	responses.BaseResponse
	RequestId string
	OrderId   string
}
