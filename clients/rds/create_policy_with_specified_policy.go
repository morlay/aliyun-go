package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreatePolicyWithSpecifiedPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PolicyId             string `position:"Query" name:"PolicyId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreatePolicyWithSpecifiedPolicyRequest) Invoke(client *sdk.Client) (resp *CreatePolicyWithSpecifiedPolicyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CreatePolicyWithSpecifiedPolicy", "rds", "")
	resp = &CreatePolicyWithSpecifiedPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreatePolicyWithSpecifiedPolicyResponse struct {
	responses.BaseResponse
	RequestId string
}
