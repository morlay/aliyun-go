package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRegionsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r DescribeRegionsRequest) Invoke(client *sdk.Client) (response *DescribeRegionsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRegionsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeRegions", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRegionsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRegionsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRegionsResponse struct {
	RequestId string
	Regions   DescribeRegionsRegionList
}

type DescribeRegionsRegion struct {
	RegionId  string
	LocalName string
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
