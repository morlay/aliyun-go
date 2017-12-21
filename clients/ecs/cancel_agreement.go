package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelAgreementRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AgreementType        string `position:"Query" name:"AgreementType"`
}

func (r CancelAgreementRequest) Invoke(client *sdk.Client) (response *CancelAgreementResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CancelAgreementRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CancelAgreement", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CancelAgreementResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CancelAgreementResponse

	err = client.DoAction(&req, &resp)
	return
}

type CancelAgreementResponse struct {
	RequestId string
}
