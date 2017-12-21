package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyInstanceAutoRenewAttributeRequest struct {
	Duration             string `position:"Query" name:"Duration"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AutoRenew            string `position:"Query" name:"AutoRenew"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyInstanceAutoRenewAttributeRequest) Invoke(client *sdk.Client) (response *ModifyInstanceAutoRenewAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceAutoRenewAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyInstanceAutoRenewAttribute", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceAutoRenewAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyInstanceAutoRenewAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceAutoRenewAttributeResponse struct {
	RequestId string
}
