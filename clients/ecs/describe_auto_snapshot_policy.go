package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAutoSnapshotPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeAutoSnapshotPolicyRequest) Invoke(client *sdk.Client) (response *DescribeAutoSnapshotPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAutoSnapshotPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeAutoSnapshotPolicy", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAutoSnapshotPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAutoSnapshotPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAutoSnapshotPolicyResponse struct {
	RequestId                  string
	AutoSnapshotOccupation     int
	AutoSnapshotPolicy         DescribeAutoSnapshotPolicyAutoSnapshotPolicy
	AutoSnapshotExcutionStatus DescribeAutoSnapshotPolicyAutoSnapshotExcutionStatus
}

type DescribeAutoSnapshotPolicyAutoSnapshotPolicy struct {
	SystemDiskPolicyEnabled           string
	SystemDiskPolicyTimePeriod        string
	SystemDiskPolicyRetentionDays     string
	SystemDiskPolicyRetentionLastWeek string
	DataDiskPolicyEnabled             string
	DataDiskPolicyTimePeriod          string
	DataDiskPolicyRetentionDays       string
	DataDiskPolicyRetentionLastWeek   string
}

type DescribeAutoSnapshotPolicyAutoSnapshotExcutionStatus struct {
	SystemDiskExcutionStatus string
	DataDiskExcutionStatus   string
}
