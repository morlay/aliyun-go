package bss

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCouponListRequest struct {
	requests.RpcRequest
	Status            string `position:"Query" name:"Status"`
	StartDeliveryTime string `position:"Query" name:"StartDeliveryTime"`
	EndDeliveryTime   string `position:"Query" name:"EndDeliveryTime"`
	PageSize          int    `position:"Query" name:"PageSize"`
	PageNum           int    `position:"Query" name:"PageNum"`
}

func (req *DescribeCouponListRequest) Invoke(client *sdk.Client) (resp *DescribeCouponListResponse, err error) {
	req.InitWithApiInfo("Bss", "2014-07-14", "DescribeCouponList", "", "")
	resp = &DescribeCouponListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCouponListResponse struct {
	responses.BaseResponse
	RequestId string
	Coupons   DescribeCouponListCouponList
}

type DescribeCouponListCoupon struct {
	CouponTemplateId int64
	TotalAmount      string
	BalanceAmount    string
	FrozenAmount     string
	ExpiredAmount    string
	DeliveryTime     string
	ExpiredTime      string
	CouponNumber     string
	Status           string
	Description      string
	CreationTime     string
	ModificationTime string
	PriceLimit       string
	Application      string
	ProductCodes     DescribeCouponListProductCodeList
	TradeTypes       DescribeCouponListTradeTypeList
}

type DescribeCouponListCouponList []DescribeCouponListCoupon

func (list *DescribeCouponListCouponList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCouponListCoupon)
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

type DescribeCouponListProductCodeList []string

func (list *DescribeCouponListProductCodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeCouponListTradeTypeList []string

func (list *DescribeCouponListTradeTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
