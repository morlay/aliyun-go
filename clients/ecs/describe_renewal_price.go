package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRenewalPriceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	Period               int    `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PriceUnit            string `position:"Query" name:"PriceUnit"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
}

func (req *DescribeRenewalPriceRequest) Invoke(client *sdk.Client) (resp *DescribeRenewalPriceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRenewalPrice", "ecs", "")
	resp = &DescribeRenewalPriceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRenewalPriceResponse struct {
	responses.BaseResponse
	RequestId common.String
	PriceInfo DescribeRenewalPricePriceInfo
}

type DescribeRenewalPricePriceInfo struct {
	Rules DescribeRenewalPriceRuleList
	Price DescribeRenewalPricePrice
}

type DescribeRenewalPriceRule struct {
	RuleId      common.Long
	Description common.String
}

type DescribeRenewalPricePrice struct {
	OriginalPrice common.Float
	DiscountPrice common.Float
	TradePrice    common.Float
	Currency      common.String
}

type DescribeRenewalPriceRuleList []DescribeRenewalPriceRule

func (list *DescribeRenewalPriceRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRenewalPriceRule)
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
