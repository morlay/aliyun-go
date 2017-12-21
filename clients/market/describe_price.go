package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePriceRequest struct {
	OrderType string `position:"Query" name:"OrderType"`
	Commodity string `position:"Query" name:"Commodity"`
}

func (r DescribePriceRequest) Invoke(client *sdk.Client) (response *DescribePriceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribePriceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Market", "2015-11-01", "DescribePrice", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribePriceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribePriceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribePriceResponse struct {
	ProductCode    string
	OriginalPrice  float32
	TradePrice     float32
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
