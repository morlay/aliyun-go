package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetCenInterRegionBandwidthLimitRequest struct {
	requests.RpcRequest
	LocalRegionId        string `position:"Query" name:"LocalRegionId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OppositeRegionId     string `position:"Query" name:"OppositeRegionId"`
	BandwidthLimit       int64  `position:"Query" name:"BandwidthLimit"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *SetCenInterRegionBandwidthLimitRequest) Invoke(client *sdk.Client) (resp *SetCenInterRegionBandwidthLimitResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "SetCenInterRegionBandwidthLimit", "cbn", "")
	resp = &SetCenInterRegionBandwidthLimitResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetCenInterRegionBandwidthLimitResponse struct {
	responses.BaseResponse
	RequestId string
}
