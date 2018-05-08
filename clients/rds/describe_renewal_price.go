package rds

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
	Quantity             int    `position:"Query" name:"Quantity"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CommodityCode        string `position:"Query" name:"CommodityCode"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	UsedTime             string `position:"Query" name:"UsedTime"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	PromotionCode        string `position:"Query" name:"PromotionCode"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	TimeType             string `position:"Query" name:"TimeType"`
	PayType              string `position:"Query" name:"PayType"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	OrderType            string `position:"Query" name:"OrderType"`
}

func (req *DescribeRenewalPriceRequest) Invoke(client *sdk.Client) (resp *DescribeRenewalPriceResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeRenewalPrice", "rds", "")
	resp = &DescribeRenewalPriceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRenewalPriceResponse struct {
	responses.BaseResponse
	RequestId common.String
	Rules     DescribeRenewalPriceRuleList
	PriceInfo DescribeRenewalPricePriceInfo
}

type DescribeRenewalPriceRule struct {
	RuleId      common.Long
	Name        common.String
	Description common.String
}

type DescribeRenewalPricePriceInfo struct {
	Currency      common.String
	OriginalPrice common.Float
	TradePrice    common.Float
	DiscountPrice common.Float
	Coupons       DescribeRenewalPriceCouponList
	RuleIds       DescribeRenewalPriceRuleIdList
	ActivityInfo  DescribeRenewalPriceActivityInfo
}

type DescribeRenewalPriceCoupon struct {
	CouponNo    common.String
	Name        common.String
	Description common.String
	IsSelected  common.String
}

type DescribeRenewalPriceActivityInfo struct {
	CheckErrMsg common.String
	ErrorCode   common.String
	Success     common.String
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

type DescribeRenewalPriceCouponList []DescribeRenewalPriceCoupon

func (list *DescribeRenewalPriceCouponList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRenewalPriceCoupon)
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

type DescribeRenewalPriceRuleIdList []common.String

func (list *DescribeRenewalPriceRuleIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
