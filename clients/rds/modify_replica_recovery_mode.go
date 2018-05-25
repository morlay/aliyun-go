package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyReplicaRecoveryModeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RecoveryMode         string `position:"Query" name:"RecoveryMode"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyReplicaRecoveryModeRequest) Invoke(client *sdk.Client) (resp *ModifyReplicaRecoveryModeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyReplicaRecoveryMode", "rds", "")
	resp = &ModifyReplicaRecoveryModeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyReplicaRecoveryModeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
