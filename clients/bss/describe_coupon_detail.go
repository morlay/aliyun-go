package bss

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCouponDetailRequest struct {
	requests.RpcRequest
	CouponNumber string `position:"Query" name:"CouponNumber"`
}

func (req *DescribeCouponDetailRequest) Invoke(client *sdk.Client) (resp *DescribeCouponDetailResponse, err error) {
	req.InitWithApiInfo("Bss", "2014-07-14", "DescribeCouponDetail", "", "")
	resp = &DescribeCouponDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCouponDetailResponse struct {
	responses.BaseResponse
	RequestId        common.String
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
	ProductCodes     DescribeCouponDetailProductCodeList
	TradeTypes       DescribeCouponDetailTradeTypeList
}

type DescribeCouponDetailProductCodeList []common.String

func (list *DescribeCouponDetailProductCodeList) UnmarshalJSON(data []byte) error {
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

type DescribeCouponDetailTradeTypeList []common.String

func (list *DescribeCouponDetailTradeTypeList) UnmarshalJSON(data []byte) error {
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
