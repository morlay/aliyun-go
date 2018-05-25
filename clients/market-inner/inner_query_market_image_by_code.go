package market_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerQueryMarketImageByCodeRequest struct {
	requests.RpcRequest
	ImagePc string `position:"Query" name:"ImagePc"`
}

func (req *InnerQueryMarketImageByCodeRequest) Invoke(client *sdk.Client) (resp *InnerQueryMarketImageByCodeResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerQueryMarketImageByCode", "yunmarket", "")
	resp = &InnerQueryMarketImageByCodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerQueryMarketImageByCodeResponse struct {
	responses.BaseResponse
	RequestId          common.String
	Success            bool
	MarketImageProduct InnerQueryMarketImageByCodeMarketImageProduct
}

type InnerQueryMarketImageByCodeMarketImageProduct struct {
	ImagePc             common.String
	ExtendInfo          common.String
	ImageRegionInfoList InnerQueryMarketImageByCodeImageRegionInfoList
	SupportChargeTypes  InnerQueryMarketImageByCodeSupportChargeTypeList
}

type InnerQueryMarketImageByCodeImageRegionInfo struct {
	RegionNo     common.String
	ImageId      common.String
	ImageVersion common.String
}

type InnerQueryMarketImageByCodeImageRegionInfoList []InnerQueryMarketImageByCodeImageRegionInfo

func (list *InnerQueryMarketImageByCodeImageRegionInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryMarketImageByCodeImageRegionInfo)
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

type InnerQueryMarketImageByCodeSupportChargeTypeList []common.String

func (list *InnerQueryMarketImageByCodeSupportChargeTypeList) UnmarshalJSON(data []byte) error {
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
