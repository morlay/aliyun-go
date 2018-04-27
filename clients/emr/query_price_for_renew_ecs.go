package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPriceForRenewEcsRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	EcsPeriod       int    `position:"Query" name:"EcsPeriod"`
	EmrPeriod       int    `position:"Query" name:"EmrPeriod"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	EcsId           string `position:"Query" name:"EcsId"`
	EcsIdList       string `position:"Query" name:"EcsIdList"`
}

func (req *QueryPriceForRenewEcsRequest) Invoke(client *sdk.Client) (resp *QueryPriceForRenewEcsResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryPriceForRenewEcs", "", "")
	resp = &QueryPriceForRenewEcsResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPriceForRenewEcsResponse struct {
	responses.BaseResponse
	RequestId  string
	EcsId      string
	EmrPrice   string
	EcsPrice   string
	EmrPriceDO QueryPriceForRenewEcsEmrPriceDO
	EcsPriceDO QueryPriceForRenewEcsEcsPriceDO
}

type QueryPriceForRenewEcsEmrPriceDO struct {
	OriginalPrice string
	DiscountPrice string
	TradePrice    string
	TaxPrice      string
	Currency      string
}

type QueryPriceForRenewEcsEcsPriceDO struct {
	OriginalPrice string
	DiscountPrice string
	TradePrice    string
	TaxPrice      string
	Currency      string
}
