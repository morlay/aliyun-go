package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	EcsId      common.String
	EmrPrice   common.String
	EcsPrice   common.String
	EmrPriceDO QueryPriceForRenewEcsEmrPriceDO
	EcsPriceDO QueryPriceForRenewEcsEcsPriceDO
}

type QueryPriceForRenewEcsEmrPriceDO struct {
	OriginalPrice common.String
	DiscountPrice common.String
	TradePrice    common.String
	TaxPrice      common.String
	Currency      common.String
}

type QueryPriceForRenewEcsEcsPriceDO struct {
	OriginalPrice common.String
	DiscountPrice common.String
	TradePrice    common.String
	TaxPrice      common.String
	Currency      common.String
}
