package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeStaticVerificationListRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken         string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	ReplicaId             string `position:"Query" name:"ReplicaId"`
	DestinationInstanceId string `position:"Query" name:"DestinationInstanceId"`
	SourceInstanceId      string `position:"Query" name:"SourceInstanceId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeStaticVerificationListRequest) Invoke(client *sdk.Client) (resp *DescribeStaticVerificationListResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeStaticVerificationList", "rds", "")
	resp = &DescribeStaticVerificationListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeStaticVerificationListResponse struct {
	responses.BaseResponse
	RequestId              common.String
	ReplicaId              common.String
	SourceInstanceId       common.String
	SourceDBNumber         common.Integer
	SourceTableNumber      common.Integer
	SourceCountNumber      common.Integer
	SourceDBSize           common.Integer
	DestinationInstanceId  common.String
	DestinationDBNumber    common.Integer
	DestinationTableNumber common.Integer
	DestinationCountNumber common.Integer
	DestinationDBSize      common.Integer
	ConsistencyPercent     common.String
	Items                  DescribeStaticVerificationListItemsItemList
}

type DescribeStaticVerificationListItemsItem struct {
	AbnormalType      common.String
	SourceDetail      common.String
	DestinationDetail common.String
}

type DescribeStaticVerificationListItemsItemList []DescribeStaticVerificationListItemsItem

func (list *DescribeStaticVerificationListItemsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStaticVerificationListItemsItem)
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
