package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ApplyAutoSnapshotPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId string `position:"Query" name:"AutoSnapshotPolicyId"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ApplyAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (response *ApplyAutoSnapshotPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ApplyAutoSnapshotPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ApplyAutoSnapshotPolicy", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ApplyAutoSnapshotPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ApplyAutoSnapshotPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type ApplyAutoSnapshotPolicyResponse struct {
	RequestId string
}
