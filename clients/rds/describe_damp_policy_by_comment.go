package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDampPolicyByCommentRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PolicyName           string `position:"Query" name:"PolicyName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDampPolicyByCommentRequest) Invoke(client *sdk.Client) (resp *DescribeDampPolicyByCommentResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDampPolicyByComment", "rds", "")
	resp = &DescribeDampPolicyByCommentResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDampPolicyByCommentResponse struct {
	responses.BaseResponse
	RequestId   string
	Policy      string
	TimeRules   string
	ActionRules string
	SourceRules string
	Handler     string
}
