package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDampPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TimeRules            string `position:"Query" name:"TimeRules"`
	ActionRules          string `position:"Query" name:"ActionRules"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	Handlers             string `position:"Query" name:"Handlers"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PolicyName           string `position:"Query" name:"PolicyName"`
	SourceRules          string `position:"Query" name:"SourceRules"`
}

func (req *ModifyDampPolicyRequest) Invoke(client *sdk.Client) (resp *ModifyDampPolicyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDampPolicy", "rds", "")
	resp = &ModifyDampPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDampPolicyResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PolicyId   common.String
	PolicyName common.String
}
