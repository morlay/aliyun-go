package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBInstanceSpecRequest struct {
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

func (r ModifyDBInstanceSpecRequest) Invoke(client *sdk.Client) (response *ModifyDBInstanceSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBInstanceSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceSpec", "dds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBInstanceSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBInstanceSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBInstanceSpecResponse struct {
	RequestId string
	OrderId   string
}
