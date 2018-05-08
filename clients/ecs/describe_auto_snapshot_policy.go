package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAutoSnapshotPolicyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (resp *DescribeAutoSnapshotPolicyResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeAutoSnapshotPolicy", "ecs", "")
	resp = &DescribeAutoSnapshotPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAutoSnapshotPolicyResponse struct {
	responses.BaseResponse
	RequestId                  common.String
	AutoSnapshotOccupation     common.Integer
	AutoSnapshotPolicy         DescribeAutoSnapshotPolicyAutoSnapshotPolicy
	AutoSnapshotExcutionStatus DescribeAutoSnapshotPolicyAutoSnapshotExcutionStatus
}

type DescribeAutoSnapshotPolicyAutoSnapshotPolicy struct {
	SystemDiskPolicyEnabled           common.String
	SystemDiskPolicyTimePeriod        common.String
	SystemDiskPolicyRetentionDays     common.String
	SystemDiskPolicyRetentionLastWeek common.String
	DataDiskPolicyEnabled             common.String
	DataDiskPolicyTimePeriod          common.String
	DataDiskPolicyRetentionDays       common.String
	DataDiskPolicyRetentionLastWeek   common.String
}

type DescribeAutoSnapshotPolicyAutoSnapshotExcutionStatus struct {
	SystemDiskExcutionStatus common.String
	DataDiskExcutionStatus   common.String
}
