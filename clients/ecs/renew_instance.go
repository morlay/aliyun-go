package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	PeriodUnit           string `position:"Query" name:"PeriodUnit"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r RenewInstanceRequest) Invoke(client *sdk.Client) (response *RenewInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RenewInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "RenewInstance", "ecs", "")

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
}
