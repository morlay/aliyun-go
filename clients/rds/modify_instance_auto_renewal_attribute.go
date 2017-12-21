package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceAutoRenewalAttributeRequest struct {
	Duration             string `position:"Query" name:"Duration"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AutoRenew            string `position:"Query" name:"AutoRenew"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyInstanceAutoRenewalAttributeRequest) Invoke(client *sdk.Client) (response *ModifyInstanceAutoRenewalAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceAutoRenewalAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyInstanceAutoRenewalAttribute", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceAutoRenewalAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyInstanceAutoRenewalAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceAutoRenewalAttributeResponse struct {
	RequestId string
}
