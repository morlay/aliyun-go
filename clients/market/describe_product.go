package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeProductRequest struct {
	requests.RpcRequest
	Code       string `position:"Query" name:"Code"`
	QueryDraft string `position:"Query" name:"QueryDraft"`
	AliUid     string `position:"Query" name:"AliUid"`
}

func (req *DescribeProductRequest) Invoke(client *sdk.Client) (resp *DescribeProductResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "DescribeProduct", "yunmarket", "")
	resp = &DescribeProductResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeProductResponse struct {
	responses.BaseResponse
	Code             common.String
	Name             common.String
	Type             common.String
	PicUrl           common.String
	Description      common.String
	ShortDescription common.String
	UseCount         common.Long
	Score            common.Float
	Status           common.String
	AuditStatus      common.String
	AuditFailMsg     common.String
	AuditTime        common.Long
	GmtCreated       common.Long
	GmtModified      common.Long
	ProductSkus      DescribeProductProductSkuList
	ProductExtras    DescribeProductProductExtraList
	ShopInfo         DescribeProductShopInfo
}

type DescribeProductProductSku struct {
	Name         common.String
	Code         common.String
	ChargeType   common.String
	Constraints  common.String
	Hidden       bool
	OrderPeriods DescribeProductOrderPeriodList
	Modules      DescribeProductModuleList
}

type DescribeProductOrderPeriod struct {
	Name       common.String
	PeriodType common.String
}

type DescribeProductModule struct {
	Id         common.String
	Name       common.String
	Code       common.String
	Properties DescribeProductPropertyList
}

type DescribeProductProperty struct {
	Name           common.String
	Key            common.String
	PropertyValues DescribeProductPropertyValueList
}

type DescribeProductPropertyValue struct {
	Value       common.String
	DisplayName common.String
	Type        common.String
}

type DescribeProductProductExtra struct {
	Key    common.String
	Values common.String
	Label  common.String
	Order  common.Integer
	Type   common.String
}

type DescribeProductShopInfo struct {
	Id         common.Long
	Name       common.String
	Emails     common.String
	WangWangs  DescribeProductWangWangList
	Telephones DescribeProductTelephoneList
}

type DescribeProductWangWang struct {
	UserName common.String
	Remark   common.String
}

type DescribeProductProductSkuList []DescribeProductProductSku

func (list *DescribeProductProductSkuList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductProductSku)
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

type DescribeProductProductExtraList []DescribeProductProductExtra

func (list *DescribeProductProductExtraList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductProductExtra)
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

type DescribeProductOrderPeriodList []DescribeProductOrderPeriod

func (list *DescribeProductOrderPeriodList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductOrderPeriod)
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

type DescribeProductModuleList []DescribeProductModule

func (list *DescribeProductModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductModule)
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

type DescribeProductPropertyList []DescribeProductProperty

func (list *DescribeProductPropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductProperty)
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

type DescribeProductPropertyValueList []DescribeProductPropertyValue

func (list *DescribeProductPropertyValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductPropertyValue)
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

type DescribeProductWangWangList []DescribeProductWangWang

func (list *DescribeProductWangWangList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductWangWang)
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

type DescribeProductTelephoneList []common.String

func (list *DescribeProductTelephoneList) UnmarshalJSON(data []byte) error {
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
