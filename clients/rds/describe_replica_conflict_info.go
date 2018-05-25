package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeReplicaConflictInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	PageSize             int    `position:"Query" name:"PageSize"`
}

func (req *DescribeReplicaConflictInfoRequest) Invoke(client *sdk.Client) (resp *DescribeReplicaConflictInfoResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeReplicaConflictInfo", "rds", "")
	resp = &DescribeReplicaConflictInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeReplicaConflictInfoResponse struct {
	responses.BaseResponse
	RequestId        common.String
	ReplicaId        common.String
	PagNumber        common.Integer
	PageRecordCount  common.Integer
	TotalRecordCount common.Integer
	Items            DescribeReplicaConflictInfoItemsItemList
}

type DescribeReplicaConflictInfoItemsItem struct {
	SourceInstanceId      common.String
	DestinationInstanceId common.String
	OccurTime             common.String
	DetailInfo            common.String
	ConfictKey            common.String
	ConfictReason         common.String
	DatabaseName          common.String
	RecoveryMode          common.String
	ConflictGtid          common.String
}

type DescribeReplicaConflictInfoItemsItemList []DescribeReplicaConflictInfoItemsItem

func (list *DescribeReplicaConflictInfoItemsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeReplicaConflictInfoItemsItem)
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
