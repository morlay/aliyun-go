package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId                     string
	TotalCount                    int
	PageNumber                    int
	PageSize                      int
	CenInterRegionBandwidthLimits DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimitList
}

type DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit struct {
	CenId            string
	LocalRegionId    string
	OppositeRegionId string
	BandwidthLimit   int64
	Status           string
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
