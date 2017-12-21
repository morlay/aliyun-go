package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelAutoSnapshotPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DiskIds              string `position:"Query" name:"DiskIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CancelAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (response *CancelAutoSnapshotPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CancelAutoSnapshotPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CancelAutoSnapshotPolicy", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CancelAutoSnapshotPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CancelAutoSnapshotPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type CancelAutoSnapshotPolicyResponse struct {
	RequestId string
}
