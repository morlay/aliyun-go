package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePolicyWithSpecifiedPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PolicyId             string `position:"Query" name:"PolicyId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreatePolicyWithSpecifiedPolicyRequest) Invoke(client *sdk.Client) (response *CreatePolicyWithSpecifiedPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreatePolicyWithSpecifiedPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CreatePolicyWithSpecifiedPolicy", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CreatePolicyWithSpecifiedPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreatePolicyWithSpecifiedPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreatePolicyWithSpecifiedPolicyResponse struct {
	RequestId string
}
