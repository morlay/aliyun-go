package phoenixsp_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindResourceInfoRequest struct {
	requests.RpcRequest
	Caller         string `position:"Query" name:"Caller"`
	InstanceIdList string `position:"Query" name:"InstanceIdList"`
	Aliuid         int64  `position:"Query" name:"Aliuid"`
	Bid            string `position:"Query" name:"Bid"`
	ResourceType   string `position:"Query" name:"ResourceType"`
}

func (req *FindResourceInfoRequest) Invoke(client *sdk.Client) (resp *FindResourceInfoResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "FindResourceInfo", "", "")
	resp = &FindResourceInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindResourceInfoResponse struct {
	responses.BaseResponse
	RequestId string
	Code      string
	Count     int
	Datas     FindResourceInfoDataList
}

type FindResourceInfoData struct {
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

type FindResourceInfoDataList []FindResourceInfoData

func (list *FindResourceInfoDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindResourceInfoData)
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
