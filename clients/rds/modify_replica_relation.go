package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyReplicaRelationRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyReplicaRelationRequest) Invoke(client *sdk.Client) (resp *ModifyReplicaRelationResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyReplicaRelation", "rds", "")
	resp = &ModifyReplicaRelationResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyReplicaRelationResponse struct {
	responses.BaseResponse
	RequestId string
}
