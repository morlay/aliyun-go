package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    string `position:"Query" name:"DBInstanceStorage"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	FromApp              string `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	OrderType            string `position:"Query" name:"OrderType"`
}

func (req *ModifyDBInstanceSpecRequest) Invoke(client *sdk.Client) (resp *ModifyDBInstanceSpecResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceSpec", "dds", "")
	resp = &ModifyDBInstanceSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBInstanceSpecResponse struct {
	responses.BaseResponse
	RequestId string
	OrderId   string
}
