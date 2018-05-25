package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeStrategyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeStrategyRequest) Invoke(client *sdk.Client) (resp *DescribeStrategyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeStrategy", "rds", "")
	resp = &DescribeStrategyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeStrategyResponse struct {
	responses.BaseResponse
	RequestId        common.String
	ReplicaId        common.String
	RecoveryMode     common.String
	VerificationMode common.String
}
