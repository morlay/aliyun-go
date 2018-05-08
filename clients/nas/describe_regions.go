package nas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRegionsRequest struct {
	requests.RpcRequest
	PageSize   int `position:"Query" name:"PageSize"`
	PageNumber int `position:"Query" name:"PageNumber"`
}

func (req *DescribeRegionsRequest) Invoke(client *sdk.Client) (resp *DescribeRegionsResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DescribeRegions", "nas", "")
	resp = &DescribeRegionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRegionsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageSize   common.Integer
	PageNumber common.Integer
	Regions    DescribeRegionsRegionList
}

type DescribeRegionsRegion struct {
	RegionId  common.String
	LocalName common.String
}

type DescribeRegionsRegionList []DescribeRegionsRegion

func (list *DescribeRegionsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRegion)
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
