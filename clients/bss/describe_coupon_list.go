package bss

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Coupons   DescribeCouponListCouponList
}

type DescribeCouponListCoupon struct {
	CouponTemplateId common.Long
	TotalAmount      common.String
	BalanceAmount    common.String
	FrozenAmount     common.String
	ExpiredAmount    common.String
	DeliveryTime     common.String
	ExpiredTime      common.String
	CouponNumber     common.String
	Status           common.String
	Description      common.String
	CreationTime     common.String
	ModificationTime common.String
	PriceLimit       common.String
	Application      common.String
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

type DescribeCouponListProductCodeList []common.String

func (list *DescribeCouponListProductCodeList) UnmarshalJSON(data []byte) error {
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

type DescribeCouponListTradeTypeList []common.String

func (list *DescribeCouponListTradeTypeList) UnmarshalJSON(data []byte) error {
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
