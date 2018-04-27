package phoenixsp_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindUserAvailableResourcesByRequest struct {
	requests.RpcRequest
	BeginEndTime       int64  `position:"Query" name:"BeginEndTime"`
	EndResCreateTime   int64  `position:"Query" name:"EndResCreateTime"`
	BillingTag         string `position:"Query" name:"BillingTag"`
	CommodityCode      string `position:"Query" name:"CommodityCode"`
	AvailableDays      int    `position:"Query" name:"AvailableDays"`
	ResourceType       string `position:"Query" name:"ResourceType"`
	BeginResCreateTime int64  `position:"Query" name:"BeginResCreateTime"`
	CurrPage           int    `position:"Query" name:"CurrPage"`
	Caller             string `position:"Query" name:"Caller"`
	AutoRenewal        string `position:"Query" name:"AutoRenewal"`
	EndEndTime         int64  `position:"Query" name:"EndEndTime"`
	InstanceIdList     string `position:"Query" name:"InstanceIdList"`
	RenewalStatus      string `position:"Query" name:"RenewalStatus"`
	HasExpiredRes      string `position:"Query" name:"HasExpiredRes"`
	PageSize           int    `position:"Query" name:"PageSize"`
	Aliuid             int64  `position:"Query" name:"Aliuid"`
	IsPrepaid          string `position:"Query" name:"IsPrepaid"`
	Bid                string `position:"Query" name:"Bid"`
	Region             string `position:"Query" name:"Region"`
}

func (req *FindUserAvailableResourcesByRequest) Invoke(client *sdk.Client) (resp *FindUserAvailableResourcesByResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "FindUserAvailableResourcesBy", "", "")
	resp = &FindUserAvailableResourcesByResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindUserAvailableResourcesByResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Count     int
	Datas     FindUserAvailableResourcesByDataList
}

type FindUserAvailableResourcesByData struct {
	Aliuid              int64
	BuyerId             int64
	InstanceId          string
	Region              string
	ResourceType        string
	ChargeType          string
	EndTime             int64
	ReleaseTime         int64
	ExtraEndTime        int64
	ResCreateTime       int64
	BillingTag          string
	CommodityCode       string
	ResourceStatus      string
	ResourceSubStatus   string
	ExpectedReleaseTime int64
	Bid                 string
	AutoRenewal         bool
	RenewalStatus       string
	RenewalDuration     int
	RenewalCycUnit      int
	SaleCycle           string
	MarketType          string
}

type FindUserAvailableResourcesByDataList []FindUserAvailableResourcesByData

func (list *FindUserAvailableResourcesByDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindUserAvailableResourcesByData)
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
