package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyAutoSnapshotPolicyExRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId   string `position:"Query" name:"AutoSnapshotPolicyId"`
	TimePoints             string `position:"Query" name:"TimePoints"`
	RetentionDays          int    `position:"Query" name:"RetentionDays"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RepeatWeekdays         string `position:"Query" name:"RepeatWeekdays"`
	AutoSnapshotPolicyName string `position:"Query" name:"AutoSnapshotPolicyName"`
}

func (req *ModifyAutoSnapshotPolicyExRequest) Invoke(client *sdk.Client) (resp *ModifyAutoSnapshotPolicyExResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyAutoSnapshotPolicyEx", "ecs", "")
	resp = &ModifyAutoSnapshotPolicyExResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyAutoSnapshotPolicyExResponse struct {
	responses.BaseResponse
	RequestId common.String
}
