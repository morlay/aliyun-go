package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRecycleBinRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
}

func (r DescribeRecycleBinRequest) Invoke(client *sdk.Client) (response *DescribeRecycleBinResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRecycleBinRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRecycleBin", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRecycleBinResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRecycleBinResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRecycleBinResponse struct {
	RequestId        string
	TotalCount       int
	RecycleBinModels DescribeRecycleBinRecycleBinModelList
}

type DescribeRecycleBinRecycleBinModel struct {
	ResourceId        string
	RegionId          string
	ResourceType      string
	CreationTime      string
	Status            string
	RelationResources DescribeRecycleBinRelationResourceList
}

type DescribeRecycleBinRelationResource struct {
	RelationResourceId   string
	RelationResourceType string
}

type DescribeRecycleBinRecycleBinModelList []DescribeRecycleBinRecycleBinModel

func (list *DescribeRecycleBinRecycleBinModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecycleBinRecycleBinModel)
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

type DescribeRecycleBinRelationResourceList []DescribeRecycleBinRelationResource

func (list *DescribeRecycleBinRelationResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecycleBinRelationResource)
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
