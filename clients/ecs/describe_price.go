package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePriceRequest struct {
	requests.RpcRequest
	DataDisk3Size           int    `position:"Query" name:"DataDisk.3.Size"`
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId                 string `position:"Query" name:"ImageId"`
	DataDisk3Category       string `position:"Query" name:"DataDisk.3.Category"`
	IoOptimized             string `position:"Query" name:"IoOptimized"`
	InternetMaxBandwidthOut int    `position:"Query" name:"InternetMaxBandwidthOut"`
	SystemDiskCategory      string `position:"Query" name:"SystemDiskCategory"`
	DataDisk4Category       string `position:"Query" name:"DataDisk.4.Category"`
	DataDisk4Size           int    `position:"Query" name:"DataDisk.4.Size"`
	PriceUnit               string `position:"Query" name:"PriceUnit"`
	InstanceType            string `position:"Query" name:"InstanceType"`
	DataDisk2Category       string `position:"Query" name:"DataDisk.2.Category"`
	DataDisk1Size           int    `position:"Query" name:"DataDisk.1.Size"`
	Period                  int    `position:"Query" name:"Period"`
	Amount                  int    `position:"Query" name:"Amount"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	DataDisk2Size           int    `position:"Query" name:"DataDisk.2.Size"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	ResourceType            string `position:"Query" name:"ResourceType"`
	DataDisk1Category       string `position:"Query" name:"DataDisk.1.Category"`
	SystemDiskSize          int    `position:"Query" name:"SystemDiskSize"`
	InternetChargeType      string `position:"Query" name:"InternetChargeType"`
	InstanceNetworkType     string `position:"Query" name:"InstanceNetworkType"`
}

func (req *DescribePriceRequest) Invoke(client *sdk.Client) (resp *DescribePriceResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribePrice", "ecs", "")
	resp = &DescribePriceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePriceResponse struct {
	responses.BaseResponse
	RequestId string
	PriceInfo DescribePricePriceInfo
}

type DescribePricePriceInfo struct {
	Rules DescribePriceRuleList
	Price DescribePricePrice
}

type DescribePriceRule struct {
	RuleId      int64
	Description string
}

type DescribePricePrice struct {
	OriginalPrice float32
	DiscountPrice float32
	TradePrice    float32
	Currency      string
}

type DescribePriceRuleList []DescribePriceRule

func (list *DescribePriceRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePriceRule)
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
