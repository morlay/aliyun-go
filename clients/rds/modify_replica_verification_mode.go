package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyReplicaVerificationModeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VerificationMode     string `position:"Query" name:"VerificationMode"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyReplicaVerificationModeRequest) Invoke(client *sdk.Client) (resp *ModifyReplicaVerificationModeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyReplicaVerificationMode", "rds", "")
	resp = &ModifyReplicaVerificationModeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyReplicaVerificationModeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
