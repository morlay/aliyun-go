package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeImageSupportInstanceTypesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeImageSupportInstanceTypesRequest) Invoke(client *sdk.Client) (response *DescribeImageSupportInstanceTypesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeImageSupportInstanceTypesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeImageSupportInstanceTypes", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeImageSupportInstanceTypesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeImageSupportInstanceTypesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeImageSupportInstanceTypesResponse struct {
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
