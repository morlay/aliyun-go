package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateDampPolicyRequest struct {
	requests.RpcRequest
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

func (req *CreateDampPolicyRequest) Invoke(client *sdk.Client) (resp *CreateDampPolicyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreateDampPolicy", "rds", "")
	resp = &CreateDampPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDampPolicyResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PolicyId   common.String
	PolicyName common.String
}
