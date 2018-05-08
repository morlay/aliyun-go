package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeResourceByTagsRequest struct {
	requests.RpcRequest
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
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

func (req *DescribeResourceByTagsRequest) Invoke(client *sdk.Client) (resp *DescribeResourceByTagsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeResourceByTags", "ecs", "")
	resp = &DescribeResourceByTagsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeResourceByTagsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageSize   common.Integer
	PageNumber common.Integer
	TotalCount common.Integer
	Resources  DescribeResourceByTagsResourceList
}

type DescribeResourceByTagsResource struct {
	ResourceId   common.String
	ResourceType common.String
	RegionId     common.String
}

type DescribeResourceByTagsResourceList []DescribeResourceByTagsResource

func (list *DescribeResourceByTagsResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeResourceByTagsResource)
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
