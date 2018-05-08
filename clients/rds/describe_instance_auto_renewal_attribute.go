package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeInstanceAutoRenewalAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	ProxyId              string `position:"Query" name:"ProxyId"`
}

func (req *DescribeInstanceAutoRenewalAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeInstanceAutoRenewalAttributeResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeInstanceAutoRenewalAttribute", "rds", "")
	resp = &DescribeInstanceAutoRenewalAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstanceAutoRenewalAttributeResponse struct {
	responses.BaseResponse
	RequestId        common.String
	PageNumber       common.Integer
	TotalRecordCount common.Integer
	PageRecordCount  common.Integer
	Items            DescribeInstanceAutoRenewalAttributeItemList
}

type DescribeInstanceAutoRenewalAttributeItem struct {
	DBInstanceId common.String
	RegionId     common.String
	Duration     common.Integer
	Status       common.String
	AutoRenew    common.String
}

type DescribeInstanceAutoRenewalAttributeItemList []DescribeInstanceAutoRenewalAttributeItem

func (list *DescribeInstanceAutoRenewalAttributeItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAutoRenewalAttributeItem)
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
