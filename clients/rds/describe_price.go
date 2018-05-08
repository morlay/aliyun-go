package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribePriceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    int    `position:"Query" name:"DBInstanceStorage"`
	Quantity             int    `position:"Query" name:"Quantity"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CommodityCode        string `position:"Query" name:"CommodityCode"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	UsedTime             string `position:"Query" name:"UsedTime"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	InstanceUsedType     int    `position:"Query" name:"InstanceUsedType"`
	Engine               string `position:"Query" name:"Engine"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	TimeType             string `position:"Query" name:"TimeType"`
	PayType              string `position:"Query" name:"PayType"`
	OrderType            string `position:"Query" name:"OrderType"`
}

func (req *DescribePriceRequest) Invoke(client *sdk.Client) (resp *DescribePriceResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribePrice", "rds", "")
	resp = &DescribePriceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePriceResponse struct {
	responses.BaseResponse
	RequestId common.String
	Rules     DescribePriceRuleList
	PriceInfo DescribePricePriceInfo
}

type DescribePriceRule struct {
	RuleId      common.Long
	Name        common.String
	Description common.String
}

type DescribePricePriceInfo struct {
	Currency      common.String
	OriginalPrice common.Float
	TradePrice    common.Float
	DiscountPrice common.Float
	Coupons       DescribePriceCouponList
	RuleIds       DescribePriceRuleIdList
	ActivityInfo  DescribePriceActivityInfo
}

type DescribePriceCoupon struct {
	CouponNo    common.String
	Name        common.String
	Description common.String
	IsSelected  common.String
}

type DescribePriceActivityInfo struct {
	CheckErrMsg common.String
	ErrorCode   common.String
	Success     common.String
}

type DescribePriceRuleList []DescribePriceRule

func (list *DescribePriceRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePriceRule)
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

type DescribePriceCouponList []DescribePriceCoupon

func (list *DescribePriceCouponList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePriceCoupon)
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

type DescribePriceRuleIdList []common.String

func (list *DescribePriceRuleIdList) UnmarshalJSON(data []byte) error {
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
