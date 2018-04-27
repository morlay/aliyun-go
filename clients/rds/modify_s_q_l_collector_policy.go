package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySQLCollectorPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	StoragePeriod        int    `position:"Query" name:"StoragePeriod"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	SQLCollectorStatus   string `position:"Query" name:"SQLCollectorStatus"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifySQLCollectorPolicyRequest) Invoke(client *sdk.Client) (resp *ModifySQLCollectorPolicyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifySQLCollectorPolicy", "rds", "")
	resp = &ModifySQLCollectorPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySQLCollectorPolicyResponse struct {
	responses.BaseResponse
	RequestId string
}
