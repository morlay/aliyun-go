package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDampPolicyByPolicyNameRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PolicyName           string `position:"Query" name:"PolicyName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDampPolicyByPolicyNameRequest) Invoke(client *sdk.Client) (response *DescribeDampPolicyByPolicyNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDampPolicyByPolicyNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDampPolicyByPolicyName", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDampPolicyByPolicyNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDampPolicyByPolicyNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDampPolicyByPolicyNameResponse struct {
	RequestId   string
	Policy      string
	TimeRules   string
	ActionRules string
	SourceRules string
	Handler     string
}
