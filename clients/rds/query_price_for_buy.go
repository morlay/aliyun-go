package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryPriceForBuyRequest struct {
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

func (req *QueryPriceForBuyRequest) Invoke(client *sdk.Client) (resp *QueryPriceForBuyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "QueryPriceForBuy", "rds", "")
	resp = &QueryPriceForBuyResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPriceForBuyResponse struct {
	responses.BaseResponse
	RequestId common.String
	Rules     QueryPriceForBuyRuleList
	PriceInfo QueryPriceForBuyPriceInfo
}

type QueryPriceForBuyRule struct {
	RuleId      common.Long
	Name        common.String
	Description common.String
}

type QueryPriceForBuyPriceInfo struct {
	Currency      common.String
	OriginalPrice common.Float
	TradePrice    common.Float
	DiscountPrice common.Float
	Coupons       QueryPriceForBuyCouponList
	RuleIds       QueryPriceForBuyRuleIdList
	ActivityInfo  QueryPriceForBuyActivityInfo
}

type QueryPriceForBuyCoupon struct {
	CouponNo    common.String
	Name        common.String
	Description common.String
	IsSelected  common.String
}

type QueryPriceForBuyActivityInfo struct {
	CheckErrMsg common.String
	ErrorCode   common.String
	Success     common.String
}

type QueryPriceForBuyRuleList []QueryPriceForBuyRule

func (list *QueryPriceForBuyRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPriceForBuyRule)
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

type QueryPriceForBuyCouponList []QueryPriceForBuyCoupon

func (list *QueryPriceForBuyCouponList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPriceForBuyCoupon)
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

type QueryPriceForBuyRuleIdList []common.String

func (list *QueryPriceForBuyRuleIdList) UnmarshalJSON(data []byte) error {
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
