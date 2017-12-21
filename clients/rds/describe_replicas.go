package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeReplicasRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeReplicasRequest) Invoke(client *sdk.Client) (response *DescribeReplicasResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeReplicasRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeReplicas", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeReplicasResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeReplicasResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeReplicasResponse struct {
	RequestId        string
	PageNumber       string
	TotalRecordCount int
	PageRecordCount  int
	Replicas         DescribeReplicasItemsList
}

type DescribeReplicasItems struct {
	ReplicaId          string
	ReplicaDescription string
	ReplicaStatus      string
	DBInstances        DescribeReplicasItems1List
}

type DescribeReplicasItems1 struct {
	DBInstanceId string
	Role         string
}

type DescribeReplicasItemsList []DescribeReplicasItems

func (list *DescribeReplicasItemsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicasItems)
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

type DescribeReplicasItems1List []DescribeReplicasItems1

func (list *DescribeReplicasItems1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicasItems1)
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
