package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAutoSnapshotPolicyRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	TimePoints             string `position:"Query" name:"TimePoints"`
	RetentionDays          int    `position:"Query" name:"RetentionDays"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RepeatWeekdays         string `position:"Query" name:"RepeatWeekdays"`
	AutoSnapshotPolicyName string `position:"Query" name:"AutoSnapshotPolicyName"`
}

func (r CreateAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (response *CreateAutoSnapshotPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateAutoSnapshotPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateAutoSnapshotPolicy", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateAutoSnapshotPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateAutoSnapshotPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateAutoSnapshotPolicyResponse struct {
	RequestId            string
	AutoSnapshotPolicyId string
}
