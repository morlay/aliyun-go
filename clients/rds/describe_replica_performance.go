package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeReplicaPerformanceRequest struct {
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

func (r DescribeReplicaPerformanceRequest) Invoke(client *sdk.Client) (response *DescribeReplicaPerformanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeReplicaPerformanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeReplicaPerformance", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeReplicaPerformanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeReplicaPerformanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeReplicaPerformanceResponse struct {
	RequestId       string
	StartTime       string
	EndTime         string
	ReplicaId       string
	PerformanceKeys DescribeReplicaPerformancePerformanceKeys
}

type DescribeReplicaPerformancePerformanceKeys struct {
	PerformanceKey DescribeReplicaPerformancePerformanceKeyItemList
}

type DescribeReplicaPerformancePerformanceKeyItem struct {
	Key               string
	Unit              string
	ValueFormat       string
	PerformanceValues DescribeReplicaPerformancePerformanceValues
}

type DescribeReplicaPerformancePerformanceValues struct {
	PerformanceValue DescribeReplicaPerformancePerformanceValueItemList
}

type DescribeReplicaPerformancePerformanceValueItem struct {
	Value string
	Date  string
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
