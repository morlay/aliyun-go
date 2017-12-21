package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDampPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
	TimeRules            string `position:"Query" name:"TimeRules"`
	ActionRules          string `position:"Query" name:"ActionRules"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	Handlers             string `position:"Query" name:"Handlers"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PolicyName           string `position:"Query" name:"PolicyName"`
	SourceRules          string `position:"Query" name:"SourceRules"`
}

func (r CreateDampPolicyRequest) Invoke(client *sdk.Client) (response *CreateDampPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateDampPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateDampPolicy", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CreateDampPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateDampPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateDampPolicyResponse struct {
	RequestId  string
	PolicyId   string
	PolicyName string
}
