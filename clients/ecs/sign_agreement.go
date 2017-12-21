package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SignAgreementRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AgreementType        string `position:"Query" name:"AgreementType"`
}

func (r SignAgreementRequest) Invoke(client *sdk.Client) (response *SignAgreementResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SignAgreementRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "SignAgreement", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		SignAgreementResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SignAgreementResponse

	err = client.DoAction(&req, &resp)
	return
}

type SignAgreementResponse struct {
	RequestId string
}
