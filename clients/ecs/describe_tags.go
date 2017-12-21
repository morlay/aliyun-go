package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTagsRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	PageSize             int    `position:"Query" name:"PageSize"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
}

func (r DescribeTagsRequest) Invoke(client *sdk.Client) (response *DescribeTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeTags", "ecs", "")

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
	Tags       DescribeTagsTagList
}

type DescribeTagsTag struct {
	TagKey            string
	TagValue          string
	ResourceTypeCount DescribeTagsResourceTypeCount
}

type DescribeTagsResourceTypeCount struct {
	Instance      int
	Disk          int
	Volume        int
	Image         int
	Snapshot      int
	Securitygroup int
}

type DescribeTagsTagList []DescribeTagsTag

func (list *DescribeTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTagsTag)
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
