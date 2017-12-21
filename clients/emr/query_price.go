package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPriceRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ZoneId                 string `position:"Query" name:"ZoneId"`
	MasterInstanceType     string `position:"Query" name:"MasterInstanceType"`
	CoreInstanceType       string `position:"Query" name:"CoreInstanceType"`
	TaskInstanceType       string `position:"Query" name:"TaskInstanceType"`
	MasterInstanceQuantity int    `position:"Query" name:"MasterInstanceQuantity"`
	CoreInstanceQuantity   int    `position:"Query" name:"CoreInstanceQuantity"`
	TaskInstanceQuantity   int    `position:"Query" name:"TaskInstanceQuantity"`
	MasterDiskType         string `position:"Query" name:"MasterDiskType"`
	CoreDiskType           string `position:"Query" name:"CoreDiskType"`
	TaskDiskType           string `position:"Query" name:"TaskDiskType"`
	MasterDiskQuantity     int    `position:"Query" name:"MasterDiskQuantity"`
	CoreDiskQuantity       int    `position:"Query" name:"CoreDiskQuantity"`
	TaskDiskQuantity       int    `position:"Query" name:"TaskDiskQuantity"`
	Duration               int    `position:"Query" name:"Duration"`
	IoOptimized            string `position:"Query" name:"IoOptimized"`
	ChargeType             string `position:"Query" name:"ChargeType"`
	NetType                string `position:"Query" name:"NetType"`
	Period                 int    `position:"Query" name:"Period"`
}

func (req *QueryPriceRequest) Invoke(client *sdk.Client) (resp *QueryPriceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryPrice", "", "")
	resp = &QueryPriceResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPriceResponse struct {
	responses.BaseResponse
	RequestId  string
	EmrPrice   string
	EcsPrice   string
	EmrPriceDO QueryPriceEmrPriceDO
	EcsPriceDO QueryPriceEcsPriceDO
}

type QueryPriceEmrPriceDO struct {
	OriginalPrice string
	DiscountPrice string
	TradePrice    string
}

type QueryPriceEcsPriceDO struct {
	OriginalPrice string
	DiscountPrice string
	TradePrice    string
}
