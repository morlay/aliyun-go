package phoenixsp_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Code      common.String
	Count     common.Integer
	Datas     FindResourceInfoDataList
}

type FindResourceInfoData struct {
	Aliuid              common.Long
	BuyerId             common.Long
	InstanceId          common.String
	Region              common.String
	ResourceType        common.String
	ChargeType          common.String
	EndTime             common.Long
	ReleaseTime         common.Long
	ExtraEndTime        common.Long
	ResCreateTime       common.Long
	BillingTag          common.String
	CommodityCode       common.String
	ResourceStatus      common.String
	ResourceSubStatus   common.String
	ExpectedReleaseTime common.Long
	Bid                 common.String
	AutoRenewal         bool
	RenewalStatus       common.String
	RenewalDuration     common.Integer
	RenewalCycUnit      common.Integer
	SaleCycle           common.String
	MarketType          common.String
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
