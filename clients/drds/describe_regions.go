package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRegionsRequest struct {
	requests.RpcRequest
}

func (req *DescribeRegionsRequest) Invoke(client *sdk.Client) (resp *DescribeRegionsResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeRegions", "", "")
	resp = &DescribeRegionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRegionsResponse struct {
	responses.BaseResponse
	RequestId   common.String
	Success     bool
	DrdsRegions DescribeRegionsDrdsRegionList
}

type DescribeRegionsDrdsRegion struct {
	RegionId           common.String
	RegionName         common.String
	ZoneId             common.String
	ZoneName           common.String
	InstanceSeriesList DescribeRegionsInstanceSeriesList
}

type DescribeRegionsInstanceSeries struct {
	SeriesName common.String
	SpecList   DescribeRegionsSpecListList
}

type DescribeRegionsDrdsRegionList []DescribeRegionsDrdsRegion

func (list *DescribeRegionsDrdsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsDrdsRegion)
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

type DescribeRegionsInstanceSeriesList []DescribeRegionsInstanceSeries

func (list *DescribeRegionsInstanceSeriesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsInstanceSeries)
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

type DescribeRegionsSpecListList []common.String

func (list *DescribeRegionsSpecListList) UnmarshalJSON(data []byte) error {
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
