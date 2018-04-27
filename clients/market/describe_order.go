package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	OrderId             int64
	AliUid              int64
	SupplierCompanyName string
	ProductCode         string
	ProductSkuCode      string
	ProductName         string
	PeriodType          string
	Quantity            int
	AccountQuantity     int64
	OrderType           string
	OrderStatus         string
	PayStatus           string
	PaidOn              int64
	CreatedOn           int64
	OriginalPrice       float32
	TotalPrice          float32
	PaymentPrice        float32
	CouponPrice         float32
	SupplierTelephones  DescribeOrderSupplierTelephoneList
	InstanceIds         DescribeOrderInstanceIdList
}

type DescribeOrderSupplierTelephoneList []string

func (list *DescribeOrderSupplierTelephoneList) UnmarshalJSON(data []byte) error {
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

type DescribeOrderInstanceIdList []string

func (list *DescribeOrderInstanceIdList) UnmarshalJSON(data []byte) error {
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
