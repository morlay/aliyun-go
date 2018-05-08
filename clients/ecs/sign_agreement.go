package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SignAgreementRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AgreementType        string `position:"Query" name:"AgreementType"`
}

func (req *SignAgreementRequest) Invoke(client *sdk.Client) (resp *SignAgreementResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "SignAgreement", "ecs", "")
	resp = &SignAgreementResponse{}
	err = client.DoAction(req, resp)
	return
}

type SignAgreementResponse struct {
	responses.BaseResponse
	RequestId common.String
}
