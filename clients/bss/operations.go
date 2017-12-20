package bss

import (
	"encoding/json"
)

func (c *BssClient) DescribeCashDetail(req *DescribeCashDetailArgs) (resp *DescribeCashDetailResponse, err error) {
	resp = &DescribeCashDetailResponse{}
	err = c.Request("DescribeCashDetail", req, resp)
	return
}

type DescribeCashDetailArgs struct {
}
type DescribeCashDetailResponse struct {
	RequestId            string
	BalanceAmount        string
	AmountOwed           string
	EnableThresholdAlert string
	MiniAlertThreshold   int64
	FrozenAmount         string
	CreditCardAmount     string
	RemmitanceAmount     string
	CreditLimit          string
	AvailableCredit      string
}

func (c *BssClient) DescribeCouponDetail(req *DescribeCouponDetailArgs) (resp *DescribeCouponDetailResponse, err error) {
	resp = &DescribeCouponDetailResponse{}
	err = c.Request("DescribeCouponDetail", req, resp)
	return
}

type DescribeCouponDetailArgs struct {
	CouponNumber string
}
type DescribeCouponDetailResponse struct {
	RequestId        string
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
	ProductCodes     DescribeCouponDetailProductCodeList
	TradeTypes       DescribeCouponDetailTradeTypeList
}

type DescribeCouponDetailProductCodeList []string

func (list *DescribeCouponDetailProductCodeList) UnmarshalJSON(data []byte) error {
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

type DescribeCouponDetailTradeTypeList []string

func (list *DescribeCouponDetailTradeTypeList) UnmarshalJSON(data []byte) error {
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

func (c *BssClient) SetResourceBusinessStatus(req *SetResourceBusinessStatusArgs) (resp *SetResourceBusinessStatusResponse, err error) {
	resp = &SetResourceBusinessStatusResponse{}
	err = c.Request("SetResourceBusinessStatus", req, resp)
	return
}

type SetResourceBusinessStatusArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	ResourceType         string
	ResourceId           string
	BusinessStatus       string
	OwnerAccount         string
}
type SetResourceBusinessStatusResponse struct {
	RequestId string
}

func (c *BssClient) DescribeCouponList(req *DescribeCouponListArgs) (resp *DescribeCouponListResponse, err error) {
	resp = &DescribeCouponListResponse{}
	err = c.Request("DescribeCouponList", req, resp)
	return
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
type DescribeCouponListArgs struct {
	Status            string
	StartDeliveryTime string
	EndDeliveryTime   string
	PageSize          int
	PageNum           int
}
type DescribeCouponListResponse struct {
	RequestId string
	Coupons   DescribeCouponListCouponList
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
