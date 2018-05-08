package phoenixsp_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Code      common.String
	Count     common.Integer
	Datas     FindBillingInfoByUserInPeriodDataList
}

type FindBillingInfoByUserInPeriodData struct {
	AliUID           common.Long
	ResourceType     common.String
	ResourceStatus   common.String
	InstanceId       common.String
	BillingTag       common.String
	ChargeType       common.String
	ResourceSource   common.String
	CommodityCode    common.String
	PropertyDetail   common.String
	Operation        common.String
	CommandStartTime common.Long
	ActualStartTime  common.Long
	ExpectStartTime  common.Long
	CommandEndTime   common.Long
	ActualEndTime    common.Long
	ExpectEndTime    common.Long
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
