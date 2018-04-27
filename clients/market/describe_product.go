package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code             string
	Name             string
	Type             string
	PicUrl           string
	Description      string
	ShortDescription string
	UseCount         int64
	Score            float32
	Status           string
	AuditStatus      string
	AuditFailMsg     string
	AuditTime        int64
	GmtCreated       int64
	GmtModified      int64
	ProductSkus      DescribeProductProductSkuList
	ProductExtras    DescribeProductProductExtraList
	ShopInfo         DescribeProductShopInfo
}

type DescribeProductProductSku struct {
	Name         string
	Code         string
	ChargeType   string
	Constraints  string
	Hidden       bool
	OrderPeriods DescribeProductOrderPeriodList
	Modules      DescribeProductModuleList
}

type DescribeProductOrderPeriod struct {
	Name       string
	PeriodType string
}

type DescribeProductModule struct {
	Id         string
	Name       string
	Code       string
	Properties DescribeProductPropertyList
}

type DescribeProductProperty struct {
	Name           string
	Key            string
	PropertyValues DescribeProductPropertyValueList
}

type DescribeProductPropertyValue struct {
	Value       string
	DisplayName string
	Type        string
}

type DescribeProductProductExtra struct {
	Key    string
	Values string
	Label  string
	Order  int
	Type   string
}

type DescribeProductShopInfo struct {
	Id         int64
	Name       string
	Emails     string
	WangWangs  DescribeProductWangWangList
	Telephones DescribeProductTelephoneList
}

type DescribeProductWangWang struct {
	UserName string
	Remark   string
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

type DescribeProductTelephoneList []string

func (list *DescribeProductTelephoneList) UnmarshalJSON(data []byte) error {
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
