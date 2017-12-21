package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddBuDBInstanceRelationRequest struct {
	requests.RpcRequest
	BusinessUnit         string `position:"Query" name:"BusinessUnit"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AddBuDBInstanceRelationRequest) Invoke(client *sdk.Client) (resp *AddBuDBInstanceRelationResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "AddBuDBInstanceRelation", "rds", "")
	resp = &AddBuDBInstanceRelationResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddBuDBInstanceRelationResponse struct {
	responses.BaseResponse
	RequestId      string
	BusinessUnit   string
	DBInstanceName string
}
