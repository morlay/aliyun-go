package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               string `position:"Query" name:"Period"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
}

func (r RenewInstanceRequest) Invoke(client *sdk.Client) (response *RenewInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RenewInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "RenewInstance", "rds", "")

	resp := struct {
		*responses.BaseResponse
		RenewInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RenewInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RenewInstanceResponse struct {
	RequestId string
	OrderId   int64
}
