package bss

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCouponDetailRequest struct {
	CouponNumber string `position:"Query" name:"CouponNumber"`
}

func (r DescribeCouponDetailRequest) Invoke(client *sdk.Client) (response *DescribeCouponDetailResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCouponDetailRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Bss", "2014-07-14", "DescribeCouponDetail", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCouponDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCouponDetailResponse

	err = client.DoAction(&req, &resp)
	return
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
