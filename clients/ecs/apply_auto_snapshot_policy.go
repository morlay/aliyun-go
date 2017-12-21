package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ApplyAutoSnapshotPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId string `position:"Query" name:"AutoSnapshotPolicyId"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ApplyAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (resp *ApplyAutoSnapshotPolicyResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ApplyAutoSnapshotPolicy", "ecs", "")
	resp = &ApplyAutoSnapshotPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ApplyAutoSnapshotPolicyResponse struct {
	responses.BaseResponse
	RequestId string
}
