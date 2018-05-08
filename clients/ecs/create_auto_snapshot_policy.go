package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateAutoSnapshotPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	TimePoints             string `position:"Query" name:"TimePoints"`
	RetentionDays          int    `position:"Query" name:"RetentionDays"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	RepeatWeekdays         string `position:"Query" name:"RepeatWeekdays"`
	AutoSnapshotPolicyName string `position:"Query" name:"AutoSnapshotPolicyName"`
}

func (req *CreateAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (resp *CreateAutoSnapshotPolicyResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateAutoSnapshotPolicy", "ecs", "")
	resp = &CreateAutoSnapshotPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateAutoSnapshotPolicyResponse struct {
	responses.BaseResponse
	RequestId            common.String
	AutoSnapshotPolicyId common.String
}
