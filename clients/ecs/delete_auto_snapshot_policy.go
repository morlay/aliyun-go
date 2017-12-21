package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAutoSnapshotPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId string `position:"Query" name:"AutoSnapshotPolicyId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (response *DeleteAutoSnapshotPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteAutoSnapshotPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteAutoSnapshotPolicy", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteAutoSnapshotPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteAutoSnapshotPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteAutoSnapshotPolicyResponse struct {
	RequestId string
}
