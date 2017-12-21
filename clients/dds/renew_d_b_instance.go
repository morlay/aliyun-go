package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewDBInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CouponNo             string `position:"Query" name:"CouponNo"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
}

func (r RenewDBInstanceRequest) Invoke(client *sdk.Client) (response *RenewDBInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RenewDBInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "RenewDBInstance", "dds", "")

	resp := struct {
		*responses.BaseResponse
		RenewDBInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RenewDBInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RenewDBInstanceResponse struct {
	RequestId string
	OrderId   string
}
