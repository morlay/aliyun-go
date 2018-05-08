package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CancelAutoSnapshotPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CancelAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (resp *CancelAutoSnapshotPolicyResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CancelAutoSnapshotPolicy", "ecs", "")
	resp = &CancelAutoSnapshotPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelAutoSnapshotPolicyResponse struct {
	responses.BaseResponse
	RequestId common.String
}
