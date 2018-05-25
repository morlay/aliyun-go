package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateStaticVerificationRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	ReplicaId             string `position:"Query" name:"ReplicaId"`
	DestinationInstanceId string `position:"Query" name:"DestinationInstanceId"`
	SourceInstanceId      string `position:"Query" name:"SourceInstanceId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateStaticVerificationRequest) Invoke(client *sdk.Client) (resp *CreateStaticVerificationResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateStaticVerification", "rds", "")
	resp = &CreateStaticVerificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateStaticVerificationResponse struct {
	responses.BaseResponse
	RequestId common.String
}
