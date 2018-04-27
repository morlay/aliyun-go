package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePriceRequest struct {
	requests.RpcRequest
	Commodity string `position:"Query" name:"Commodity"`
	OrderType string `position:"Query" name:"OrderType"`
}

func (req *DescribePriceRequest) Invoke(client *sdk.Client) (resp *DescribePriceResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "DescribePrice", "yunmarket", "")
	resp = &DescribePriceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePriceResponse struct {
	responses.BaseResponse
	ProductCode    string
	OriginalPrice  float32
	TradePrice     float32
	DiscountPrice  float32
	PromotionRules DescribePricePromotionRuleList
}

type DescribePricePromotionRule struct {
	RuleId string
	Name   string
	Title  string
}

type DescribePricePromotionRuleList []DescribePricePromotionRule

func (list *DescribePricePromotionRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePricePromotionRule)
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
