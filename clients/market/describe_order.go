package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeOrderRequest struct {
	requests.RpcRequest
	OrderId string `position:"Query" name:"OrderId"`
}

func (req *DescribeOrderRequest) Invoke(client *sdk.Client) (resp *DescribeOrderResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "DescribeOrder", "yunmarket", "")
	resp = &DescribeOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeOrderResponse struct {
	responses.BaseResponse
	OrderId             common.Long
	AliUid              common.Long
	SupplierCompanyName common.String
	ProductCode         common.String
	ProductSkuCode      common.String
	ProductName         common.String
	PeriodType          common.String
	Quantity            common.Integer
	AccountQuantity     common.Long
	OrderType           common.String
	OrderStatus         common.String
	PayStatus           common.String
	PaidOn              common.Long
	CreatedOn           common.Long
	OriginalPrice       common.Float
	TotalPrice          common.Float
	PaymentPrice        common.Float
	CouponPrice         common.Float
	SupplierTelephones  DescribeOrderSupplierTelephoneList
	InstanceIds         DescribeOrderInstanceIdList
}

type DescribeOrderSupplierTelephoneList []common.String

func (list *DescribeOrderSupplierTelephoneList) UnmarshalJSON(data []byte) error {
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

type DescribeOrderInstanceIdList []common.String

func (list *DescribeOrderInstanceIdList) UnmarshalJSON(data []byte) error {
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
