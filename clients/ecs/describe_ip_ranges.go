package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeIpRangesRequest struct {
	requests.RpcRequest
	NicType              string `position:"Query" name:"NicType"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ClusterId            string `position:"Query" name:"ClusterId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeIpRangesRequest) Invoke(client *sdk.Client) (resp *DescribeIpRangesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeIpRanges", "ecs", "")
	resp = &DescribeIpRangesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeIpRangesResponse struct {
	responses.BaseResponse
	RequestId  common.String
	RegionId   common.String
	ClusterId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	IpRanges   DescribeIpRangesIpRangeList
}

type DescribeIpRangesIpRange struct {
	IpAddress common.String
	NicType   common.String
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
