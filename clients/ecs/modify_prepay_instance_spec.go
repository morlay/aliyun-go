package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyPrepayInstanceSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyPrepayInstanceSpecRequest) Invoke(client *sdk.Client) (response *ModifyPrepayInstanceSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyPrepayInstanceSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyPrepayInstanceSpec", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyPrepayInstanceSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyPrepayInstanceSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyPrepayInstanceSpecResponse struct {
	RequestId string
	OrderId   string
}
