package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRenewalPriceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	Period               int    `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PriceUnit            string `position:"Query" name:"PriceUnit"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
}

func (r DescribeRenewalPriceRequest) Invoke(client *sdk.Client) (response *DescribeRenewalPriceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRenewalPriceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRenewalPrice", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRenewalPriceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRenewalPriceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRenewalPriceResponse struct {
	RequestId string
	PriceInfo DescribeRenewalPricePriceInfo
}

type DescribeRenewalPricePriceInfo struct {
	Rules DescribeRenewalPriceRuleList
	Price DescribeRenewalPricePrice
}

type DescribeRenewalPriceRule struct {
	RuleId      int64
	Description string
}

type DescribeRenewalPricePrice struct {
	OriginalPrice float32
	DiscountPrice float32
	TradePrice    float32
	Currency      string
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
