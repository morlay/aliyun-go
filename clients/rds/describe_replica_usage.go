package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeReplicaUsageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SourceDBInstanceId   string `position:"Query" name:"SourceDBInstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeReplicaUsageRequest) Invoke(client *sdk.Client) (resp *DescribeReplicaUsageResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeReplicaUsage", "rds", "")
	resp = &DescribeReplicaUsageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeReplicaUsageResponse struct {
	responses.BaseResponse
	RequestId       common.String
	StartTime       common.String
	EndTime         common.String
	ReplicaId       common.String
	PerformanceKeys DescribeReplicaUsagePerformanceKeys
}

type DescribeReplicaUsagePerformanceKeys struct {
	PerformanceKey DescribeReplicaUsagePerformanceKeyItemList
}

type DescribeReplicaUsagePerformanceKeyItem struct {
	Key               common.String
	Unit              common.String
	ValueFormat       common.String
	PerformanceValues DescribeReplicaUsagePerformanceValues
}

type DescribeReplicaUsagePerformanceValues struct {
	PerformanceValue DescribeReplicaUsagePerformanceValueItemList
}

type DescribeReplicaUsagePerformanceValueItem struct {
	Value common.String
	Date  common.String
}

type DescribeReplicaUsagePerformanceKeyItemList []DescribeReplicaUsagePerformanceKeyItem

func (list *DescribeReplicaUsagePerformanceKeyItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaUsagePerformanceKeyItem)
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

type DescribeReplicaUsagePerformanceValueItemList []DescribeReplicaUsagePerformanceValueItem

func (list *DescribeReplicaUsagePerformanceValueItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaUsagePerformanceValueItem)
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
