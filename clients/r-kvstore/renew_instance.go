package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RenewInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	FromApp              string `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceClass        string `position:"Query" name:"InstanceClass"`
	Capacity             string `position:"Query" name:"Capacity"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ForceUpgrade         string `position:"Query" name:"ForceUpgrade"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
}

func (req *RenewInstanceRequest) Invoke(client *sdk.Client) (resp *RenewInstanceResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "RenewInstance", "redisa", "")
	resp = &RenewInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
	OrderId   common.String
	EndTime   common.String
}
