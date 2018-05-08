package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeTagsRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	DistinctKey          string `position:"Query" name:"DistinctKey"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeTagsRequest) Invoke(client *sdk.Client) (resp *DescribeTagsResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeTags", "slb", "")
	resp = &DescribeTagsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTagsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageSize   common.Integer
	PageNumber common.Integer
	TotalCount common.Integer
	TagSets    DescribeTagsTagSetList
}

type DescribeTagsTagSet struct {
	TagKey        common.String
	TagValue      common.String
	InstanceCount common.Integer
}

type DescribeTagsTagSetList []DescribeTagsTagSet

func (list *DescribeTagsTagSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTagsTagSet)
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
