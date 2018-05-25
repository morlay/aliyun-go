package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateOrderRequest struct {
	requests.RpcRequest
	OrderSouce  string `position:"Query" name:"OrderSouce"`
	Commodity   string `position:"Query" name:"Commodity"`
	ClientToken string `position:"Query" name:"ClientToken"`
	OwnerId     string `position:"Query" name:"OwnerId"`
	PaymentType string `position:"Query" name:"PaymentType"`
	OrderType   string `position:"Query" name:"OrderType"`
}

func (req *CreateOrderRequest) Invoke(client *sdk.Client) (resp *CreateOrderResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "CreateOrder", "yunmarket", "")
	resp = &CreateOrderResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateOrderResponse struct {
	responses.BaseResponse
	RequestId   common.String
	OrderId     common.String
	InstanceIds CreateOrderInstanceIdList
}

type CreateOrderInstanceIdList []common.String

func (list *CreateOrderInstanceIdList) UnmarshalJSON(data []byte) error {
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
