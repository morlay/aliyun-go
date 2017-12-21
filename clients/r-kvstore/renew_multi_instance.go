package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewMultiInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int64  `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	FromApp              string `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
}

func (req *RenewMultiInstanceRequest) Invoke(client *sdk.Client) (resp *RenewMultiInstanceResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "RenewMultiInstance", "redisa", "")
	resp = &RenewMultiInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewMultiInstanceResponse struct {
	responses.BaseResponse
	RequestId string
	OrderId   string
}
