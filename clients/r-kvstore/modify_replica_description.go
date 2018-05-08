package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyReplicaDescriptionRequest struct {
	requests.RpcRequest
	ReplicaDescription   string `position:"Query" name:"ReplicaDescription"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyReplicaDescriptionRequest) Invoke(client *sdk.Client) (resp *ModifyReplicaDescriptionResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyReplicaDescription", "redisa", "")
	resp = &ModifyReplicaDescriptionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyReplicaDescriptionResponse struct {
	responses.BaseResponse
	RequestId common.String
}
