package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCenInterRegionBandwidthLimitsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeCenInterRegionBandwidthLimitsRequest) Invoke(client *sdk.Client) (resp *DescribeCenInterRegionBandwidthLimitsResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenInterRegionBandwidthLimits", "cbn", "")
	resp = &DescribeCenInterRegionBandwidthLimitsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCenInterRegionBandwidthLimitsResponse struct {
	responses.BaseResponse
	RequestId                     common.String
	TotalCount                    common.Integer
	PageNumber                    common.Integer
	PageSize                      common.Integer
	CenInterRegionBandwidthLimits DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimitList
}

type DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit struct {
	CenId            common.String
	LocalRegionId    common.String
	OppositeRegionId common.String
	BandwidthLimit   common.Long
	Status           common.String
}

type DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimitList []DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit

func (list *DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimitList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit)
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
