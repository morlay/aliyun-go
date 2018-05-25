package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVerificationListRequest struct {
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

func (req *DescribeVerificationListRequest) Invoke(client *sdk.Client) (resp *DescribeVerificationListResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeVerificationList", "rds", "")
	resp = &DescribeVerificationListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVerificationListResponse struct {
	responses.BaseResponse
	RequestId        common.String
	ReplicaId        common.String
	PagNumber        common.Integer
	PageRecordCount  common.Integer
	TotalRecordCount common.Integer
	Items            DescribeVerificationListItemsItemList
}

type DescribeVerificationListItemsItem struct {
	InstanceIdA        common.String
	InstanceIdB        common.String
	Key                common.String
	KeyType            common.String
	InconsistentType   common.String
	OccurTime          common.String
	Schema             common.String
	InconsistentFields common.String
}

type DescribeVerificationListItemsItemList []DescribeVerificationListItemsItem

func (list *DescribeVerificationListItemsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVerificationListItemsItem)
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
