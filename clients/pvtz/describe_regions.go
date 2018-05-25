package pvtz

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRegionsRequest struct {
	requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *DescribeRegionsRequest) Invoke(client *sdk.Client) (resp *DescribeRegionsResponse, err error) {
	req.InitWithApiInfo("pvtz", "2018-01-01", "DescribeRegions", "pvtz", "")
	resp = &DescribeRegionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRegionsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Regions   DescribeRegionsRegionList
}

type DescribeRegionsRegion struct {
	RegionId   common.String
	RegionName common.String
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
