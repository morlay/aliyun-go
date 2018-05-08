package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeAutoSnapshotPolicyExRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AutoSnapshotPolicyId string `position:"Query" name:"AutoSnapshotPolicyId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeAutoSnapshotPolicyExRequest) Invoke(client *sdk.Client) (resp *DescribeAutoSnapshotPolicyExResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeAutoSnapshotPolicyEx", "ecs", "")
	resp = &DescribeAutoSnapshotPolicyExResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAutoSnapshotPolicyExResponse struct {
	responses.BaseResponse
	RequestId            common.String
	TotalCount           common.Integer
	PageNumber           common.Integer
	PageSize             common.Integer
	AutoSnapshotPolicies DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList
}

type DescribeAutoSnapshotPolicyExAutoSnapshotPolicy struct {
	AutoSnapshotPolicyId   common.String
	RegionId               common.String
	AutoSnapshotPolicyName common.String
	TimePoints             common.String
	RepeatWeekdays         common.String
	RetentionDays          common.Integer
	DiskNums               common.Integer
	VolumeNums             common.Integer
	CreationTime           common.String
	Status                 common.String
}

type DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList []DescribeAutoSnapshotPolicyExAutoSnapshotPolicy

func (list *DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAutoSnapshotPolicyExAutoSnapshotPolicy)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
