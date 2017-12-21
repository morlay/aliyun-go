package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeIpRangesRequest struct {
	NicType              string `position:"Query" name:"NicType"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ClusterId            string `position:"Query" name:"ClusterId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeIpRangesRequest) Invoke(client *sdk.Client) (response *DescribeIpRangesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeIpRangesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeIpRanges", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeIpRangesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeIpRangesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeIpRangesResponse struct {
	RequestId  string
	RegionId   string
	ClusterId  string
	TotalCount int
	PageNumber int
	PageSize   int
	IpRanges   DescribeIpRangesIpRangeList
}

type DescribeIpRangesIpRange struct {
	IpAddress string
	NicType   string
}

type DescribeIpRangesIpRangeList []DescribeIpRangesIpRange

func (list *DescribeIpRangesIpRangeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeIpRangesIpRange)
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
