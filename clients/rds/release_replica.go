package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReleaseReplicaRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReleaseReplicaRequest) Invoke(client *sdk.Client) (resp *ReleaseReplicaResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ReleaseReplica", "rds", "")
	resp = &ReleaseReplicaResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseReplicaResponse struct {
	responses.BaseResponse
	RequestId common.String
}
