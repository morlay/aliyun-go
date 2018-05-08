package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeReplicasRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeReplicasRequest) Invoke(client *sdk.Client) (resp *DescribeReplicasResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeReplicas", "rds", "")
	resp = &DescribeReplicasResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeReplicasResponse struct {
	responses.BaseResponse
	RequestId        common.String
	PageNumber       common.String
	TotalRecordCount common.Integer
	PageRecordCount  common.Integer
	Replicas         DescribeReplicasItemsList
}

type DescribeReplicasItems struct {
	ReplicaId          common.String
	ReplicaDescription common.String
	ReplicaStatus      common.String
	ReplicaMode        common.String
	DomainMode         common.String
	DBInstances        DescribeReplicasItems1List
}

type DescribeReplicasItems1 struct {
	DBInstanceId  common.String
	Role          common.String
	ReadWriteType common.String
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
