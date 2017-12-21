package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSQLCollectorPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeSQLCollectorPolicyRequest) Invoke(client *sdk.Client) (resp *DescribeSQLCollectorPolicyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLCollectorPolicy", "rds", "")
	resp = &DescribeSQLCollectorPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSQLCollectorPolicyResponse struct {
	responses.BaseResponse
	RequestId          string
	SQLCollectorStatus string
}
