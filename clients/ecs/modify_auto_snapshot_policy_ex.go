package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyAutoSnapshotPolicyExRequest struct {
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId   string `position:"Query" name:"AutoSnapshotPolicyId"`
	TimePoints             string `position:"Query" name:"TimePoints"`
	RetentionDays          int    `position:"Query" name:"RetentionDays"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RepeatWeekdays         string `position:"Query" name:"RepeatWeekdays"`
	AutoSnapshotPolicyName string `position:"Query" name:"AutoSnapshotPolicyName"`
}

func (r ModifyAutoSnapshotPolicyExRequest) Invoke(client *sdk.Client) (response *ModifyAutoSnapshotPolicyExResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyAutoSnapshotPolicyExRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyAutoSnapshotPolicyEx", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyAutoSnapshotPolicyExResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyAutoSnapshotPolicyExResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyAutoSnapshotPolicyExResponse struct {
	RequestId string
}
