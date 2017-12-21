package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAutoSnapshotPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId string `position:"Query" name:"AutoSnapshotPolicyId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (resp *DeleteAutoSnapshotPolicyResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteAutoSnapshotPolicy", "ecs", "")
	resp = &DeleteAutoSnapshotPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAutoSnapshotPolicyResponse struct {
	responses.BaseResponse
	RequestId string
}
