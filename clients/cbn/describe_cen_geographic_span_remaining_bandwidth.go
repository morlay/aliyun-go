package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCenGeographicSpanRemainingBandwidthRequest struct {
	requests.RpcRequest
	GeographicRegionBId  string `position:"Query" name:"GeographicRegionBId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	GeographicRegionAId  string `position:"Query" name:"GeographicRegionAId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeCenGeographicSpanRemainingBandwidthRequest) Invoke(client *sdk.Client) (resp *DescribeCenGeographicSpanRemainingBandwidthResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenGeographicSpanRemainingBandwidth", "cbn", "")
	resp = &DescribeCenGeographicSpanRemainingBandwidthResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCenGeographicSpanRemainingBandwidthResponse struct {
	responses.BaseResponse
	RequestId          common.String
	RemainingBandwidth common.Long
}
