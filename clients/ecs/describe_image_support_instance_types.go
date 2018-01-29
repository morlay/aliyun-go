package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeImageSupportInstanceTypesRequest struct {
	requests.RpcRequest
	ActionType           string                                       `position:"Query" name:"ActionType"`
	Filters              *DescribeImageSupportInstanceTypesFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                        `position:"Query" name:"ResourceOwnerId"`
	ImageId              string                                       `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string                                       `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64                                        `position:"Query" name:"OwnerId"`
}

func (req *DescribeImageSupportInstanceTypesRequest) Invoke(client *sdk.Client) (resp *DescribeImageSupportInstanceTypesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeImageSupportInstanceTypes", "ecs", "")
	resp = &DescribeImageSupportInstanceTypesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeImageSupportInstanceTypesFilter struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type DescribeImageSupportInstanceTypesResponse struct {
	responses.BaseResponse
	RequestId     string
	RegionId      string
	ImageId       string
	InstanceTypes DescribeImageSupportInstanceTypesInstanceTypeList
}

type DescribeImageSupportInstanceTypesInstanceType struct {
	InstanceTypeId     string
	CpuCoreCount       int
	MemorySize         float32
	InstanceTypeFamily string
}

type DescribeImageSupportInstanceTypesFilterList []DescribeImageSupportInstanceTypesFilter

func (list *DescribeImageSupportInstanceTypesFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSupportInstanceTypesFilter)
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

type DescribeImageSupportInstanceTypesInstanceTypeList []DescribeImageSupportInstanceTypesInstanceType

func (list *DescribeImageSupportInstanceTypesInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSupportInstanceTypesInstanceType)
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
