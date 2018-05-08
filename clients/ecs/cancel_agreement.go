package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CancelAgreementRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AgreementType        string `position:"Query" name:"AgreementType"`
}

func (req *CancelAgreementRequest) Invoke(client *sdk.Client) (resp *CancelAgreementResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CancelAgreement", "ecs", "")
	resp = &CancelAgreementResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelAgreementResponse struct {
	responses.BaseResponse
	RequestId common.String
}
