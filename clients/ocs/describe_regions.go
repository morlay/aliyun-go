package ocs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRegionsRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (r DescribeRegionsRequest) Invoke(client *sdk.Client) (response *DescribeRegionsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRegionsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "DescribeRegions", "", "")

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
	RequestId                  string
	GetDescribeRegionsResponse DescribeRegionsGetDescribeRegionsResponse
}

type DescribeRegionsGetDescribeRegionsResponse struct {
	Total    int
	PageNo   int
	PageSize int
	Regions  DescribeRegionsRegionList
}

type DescribeRegionsRegion struct {
	RegionId  string
	IzNumber  string
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
