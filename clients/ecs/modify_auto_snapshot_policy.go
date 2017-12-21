package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyAutoSnapshotPolicyRequest struct {
	requests.RpcRequest
	DataDiskPolicyEnabled             string `position:"Query" name:"DataDiskPolicyEnabled"`
	ResourceOwnerId                   int64  `position:"Query" name:"ResourceOwnerId"`
	DataDiskPolicyRetentionDays       int    `position:"Query" name:"DataDiskPolicyRetentionDays"`
	ResourceOwnerAccount              string `position:"Query" name:"ResourceOwnerAccount"`
	SystemDiskPolicyRetentionLastWeek string `position:"Query" name:"SystemDiskPolicyRetentionLastWeek"`
	OwnerAccount                      string `position:"Query" name:"OwnerAccount"`
	SystemDiskPolicyTimePeriod        int    `position:"Query" name:"SystemDiskPolicyTimePeriod"`
	OwnerId                           int64  `position:"Query" name:"OwnerId"`
	DataDiskPolicyRetentionLastWeek   string `position:"Query" name:"DataDiskPolicyRetentionLastWeek"`
	SystemDiskPolicyRetentionDays     int    `position:"Query" name:"SystemDiskPolicyRetentionDays"`
	DataDiskPolicyTimePeriod          int    `position:"Query" name:"DataDiskPolicyTimePeriod"`
	SystemDiskPolicyEnabled           string `position:"Query" name:"SystemDiskPolicyEnabled"`
}

func (req *ModifyAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (resp *ModifyAutoSnapshotPolicyResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyAutoSnapshotPolicy", "ecs", "")
	resp = &ModifyAutoSnapshotPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyAutoSnapshotPolicyResponse struct {
	responses.BaseResponse
	RequestId string
}
