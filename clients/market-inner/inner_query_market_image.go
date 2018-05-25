package market_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerQueryMarketImageRequest struct {
	requests.RpcRequest
	ImagePc string `position:"Query" name:"ImagePc"`
	AliUid  int64  `position:"Query" name:"AliUid"`
}

func (req *InnerQueryMarketImageRequest) Invoke(client *sdk.Client) (resp *InnerQueryMarketImageResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerQueryMarketImage", "yunmarket", "")
	resp = &InnerQueryMarketImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerQueryMarketImageResponse struct {
	responses.BaseResponse
	RequestId          common.String
	Success            bool
	MarketImageProduct InnerQueryMarketImageMarketImageProduct
}

type InnerQueryMarketImageMarketImageProduct struct {
	ImagePc             common.String
	ExtendInfo          common.String
	ImageRegionInfoList InnerQueryMarketImageImageRegionInfoList
	SupportChargeTypes  InnerQueryMarketImageSupportChargeTypeList
}

type InnerQueryMarketImageImageRegionInfo struct {
	RegionNo     common.String
	ImageId      common.String
	ImageVersion common.String
}

type InnerQueryMarketImageImageRegionInfoList []InnerQueryMarketImageImageRegionInfo

func (list *InnerQueryMarketImageImageRegionInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryMarketImageImageRegionInfo)
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

type InnerQueryMarketImageSupportChargeTypeList []common.String

func (list *InnerQueryMarketImageSupportChargeTypeList) UnmarshalJSON(data []byte) error {
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
