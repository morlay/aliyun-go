package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyReplicaModeRequest struct {
	requests.RpcRequest
	DomainMode           string `position:"Query" name:"DomainMode"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PrimaryInstanceId    string `position:"Query" name:"PrimaryInstanceId"`
	ReplicaMode          string `position:"Query" name:"ReplicaMode"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyReplicaModeRequest) Invoke(client *sdk.Client) (resp *ModifyReplicaModeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyReplicaMode", "rds", "")
	resp = &ModifyReplicaModeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyReplicaModeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
