package phoenixsp_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindBillingInfoByUserInPeriodRequest struct {
	requests.RpcRequest
	CurrPage        int    `position:"Query" name:"CurrPage"`
	PageSize        int    `position:"Query" name:"PageSize"`
	EndTime         int64  `position:"Query" name:"EndTime"`
	AliUID          int64  `position:"Query" name:"AliUID"`
	CommodityCode   string `position:"Query" name:"CommodityCode"`
	StartTime       int64  `position:"Query" name:"StartTime"`
	Bid             string `position:"Query" name:"Bid"`
	TradeInstanceID string `position:"Query" name:"TradeInstanceID"`
}

func (req *FindBillingInfoByUserInPeriodRequest) Invoke(client *sdk.Client) (resp *FindBillingInfoByUserInPeriodResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "FindBillingInfoByUserInPeriod", "", "")
	resp = &FindBillingInfoByUserInPeriodResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindBillingInfoByUserInPeriodResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Count     int
	Datas     FindBillingInfoByUserInPeriodDataList
}

type FindBillingInfoByUserInPeriodData struct {
	AliUID           int64
	ResourceType     string
	ResourceStatus   string
	InstanceId       string
	BillingTag       string
	ChargeType       string
	ResourceSource   string
	CommodityCode    string
	PropertyDetail   string
	Operation        string
	CommandStartTime int64
	ActualStartTime  int64
	ExpectStartTime  int64
	CommandEndTime   int64
	ActualEndTime    int64
	ExpectEndTime    int64
}

type FindBillingInfoByUserInPeriodDataList []FindBillingInfoByUserInPeriodData

func (list *FindBillingInfoByUserInPeriodDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindBillingInfoByUserInPeriodData)
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
