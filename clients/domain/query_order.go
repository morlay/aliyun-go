package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryOrderRequest struct {
	OrderID string `position:"Query" name:"OrderID"`
}

func (r QueryOrderRequest) Invoke(client *sdk.Client) (response *QueryOrderResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryOrderRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "QueryOrder", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryOrderResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryOrderResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryOrderResponse struct {
	RequestId     string
	OrderID       string
	UserID        string
	Money         string
	OrderDate     string
	CheckFlag     bool
	CheckDate     string
	ValidFlag     bool
	CheckType     bool
	OrderProducts QueryOrderSubOrderResultList
}

type QueryOrderSubOrderResult struct {
	TrackID      string
	OrderID      string
	SaleID       string
	UserID       string
	ClassID      string
	ProductName  string
	RelatedName  string
	ActionType   string
	PeriodUnit   int
	PeriodNum    int
	Amount       int
	OrderDate    string
	CheckFlag    bool
	CheckDate    string
	BizStatus    int
	UpdateDate   string
	DeadDate     string
	ValidFlag    bool
	Money        string
	ParentSaleID string
}

type QueryOrderSubOrderResultList []QueryOrderSubOrderResult

func (list *QueryOrderSubOrderResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryOrderSubOrderResult)
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
