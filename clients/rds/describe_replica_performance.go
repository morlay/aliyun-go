package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeReplicaPerformanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SourceDBInstanceId   string `position:"Query" name:"SourceDBInstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	Key                  string `position:"Query" name:"Key"`
}

func (req *DescribeReplicaPerformanceRequest) Invoke(client *sdk.Client) (resp *DescribeReplicaPerformanceResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeReplicaPerformance", "rds", "")
	resp = &DescribeReplicaPerformanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeReplicaPerformanceResponse struct {
	responses.BaseResponse
	RequestId       common.String
	StartTime       common.String
	EndTime         common.String
	ReplicaId       common.String
	PerformanceKeys DescribeReplicaPerformancePerformanceKeys
}

type DescribeReplicaPerformancePerformanceKeys struct {
	PerformanceKey DescribeReplicaPerformancePerformanceKeyItemList
}

type DescribeReplicaPerformancePerformanceKeyItem struct {
	Key               common.String
	Unit              common.String
	ValueFormat       common.String
	PerformanceValues DescribeReplicaPerformancePerformanceValues
}

type DescribeReplicaPerformancePerformanceValues struct {
	PerformanceValue DescribeReplicaPerformancePerformanceValueItemList
}

type DescribeReplicaPerformancePerformanceValueItem struct {
	Value common.String
	Date  common.String
}

type DescribeReplicaPerformancePerformanceKeyItemList []DescribeReplicaPerformancePerformanceKeyItem

func (list *DescribeReplicaPerformancePerformanceKeyItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaPerformancePerformanceKeyItem)
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

type DescribeReplicaPerformancePerformanceValueItemList []DescribeReplicaPerformancePerformanceValueItem

func (list *DescribeReplicaPerformancePerformanceValueItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaPerformancePerformanceValueItem)
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
