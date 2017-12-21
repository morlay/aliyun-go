package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTagsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DistinctKey          string `position:"Query" name:"DistinctKey"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	PageSize             int    `position:"Query" name:"PageSize"`
}

func (r DescribeTagsRequest) Invoke(client *sdk.Client) (response *DescribeTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeTags", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeTagsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeTagsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTagsResponse struct {
	RequestId  string
	PageSize   int
	PageNumber int
	TotalCount int
	TagSets    DescribeTagsTagSetList
}

type DescribeTagsTagSet struct {
	TagKey        string
	TagValue      string
	InstanceCount int
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
